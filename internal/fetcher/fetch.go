package fetcher

import (
	"fmt"
	"sync"
)

// Fetch() retrieves data from multiple URLs concurrently and returns the aggregated results.
// It uses an owner-consumer pattern to manage concurrency:
// - The owner spawns goroutines to fetch data from each URL.
// - The consumer collects the results and handles errors.
// Returns:
//   - A slice of maps containing the fetched data.
//   - An error if one or more requests fail.
func (f *Fetcher) Fetch() ([]map[string]interface{}, error) {
	dataChan, errChan := f.owner()
	data, err := f.consumer(dataChan, errChan)

	if err != nil {
		return nil, err
	}

	return data, nil
}

// owner() initializes and returns dataChan and errChan,
// synchronizes goroutines that fetch and parse data,
// and closes channels upon completion of all goroutines
func (f *Fetcher) owner() (<-chan []map[string]interface{}, <-chan error) {
	// Create buffered channels for data and errors
	// The buffer size is set to the number of URLs to avoid blocking
	dataChan := make(chan []map[string]interface{}, len(f.Urls))
	errChan := make(chan error, len(f.Urls))

	// Use a WaitGroup to synchronize all goroutines
	wg := &sync.WaitGroup{}

	for _, url := range f.Urls {
		// Increment the WaitGroup
		wg.Add(1)

		go func(url string) {
			// Defer decrementing the WaitGroup
			defer wg.Done()
			// Fetch and parse the data from the URL
			batch, err := f.RequesterParser.RequestAndParse(url)

			if err != nil {
				// In case of error log the error and send to errChan
				f.ErrorLog.Printf("\t\nurl: %s\t\nerr: %+v", url, err)
				errChan <- err
				return
			}
			// Send data to dataChan
			dataChan <- batch
		}(url)
	}

	// Wait for all goroutines to complete and close channels
	go func() {
		wg.Wait()
		defer close(dataChan)
		defer close(errChan)
	}()

	return dataChan, errChan
}

// consumer() processes data and error messages from the provided channels
func (f *Fetcher) consumer(dataChan <-chan []map[string]interface{}, errChan <-chan error) ([]map[string]interface{}, error) {
	// Collect results from the channels
	var data []map[string]interface{}
	var errors []error

	// Process all results by looping over the number of URLs
	for i := 0; i < len(f.Urls); i++ {
		select {
		case err := <-errChan:
			errors = append(errors, err)
		case batch := <-dataChan:
			data = append(data, batch...)
		}
	}

	// Check if there were any errors
	if len(errors) > 0 {
		// Aggregate all errors and return
		return data, fmt.Errorf("one or more requests failed: %+v", errors)
	}

	return data, nil
}
