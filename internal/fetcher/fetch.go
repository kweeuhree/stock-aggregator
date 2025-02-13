package fetcher

import (
	"fmt"
	"sync"
)

func (f *Fetcher) Fetch() ([]map[string]interface{}, error) {
	wg := sync.WaitGroup{}

	dataChan := make(chan []map[string]interface{}, len(f.Urls))
	errChan := make(chan error, len(f.Urls))

	for _, url := range f.Urls {
		wg.Add(1)

		go func(url string) {
			defer wg.Done()
			batch, err := f.RequesterParser.RequestAndParse(url)

			if err != nil {
				f.ErrorLog.Printf("\t\nurl: %s\t\nerr: %+v", url, err)
				errChan <- err
				return
			}

			dataChan <- batch
		}(url)
	}

	// Wait for all goroutines to finish
	go func() {
		wg.Wait()
		close(dataChan)
		close(errChan)
	}()

	// Collect results from the channels
	var data []map[string]interface{}
	var errors []error

	// Collect errors from errChan
	for err := range errChan {
		errors = append(errors, err)
	}

	// Check if there were any errors
	if len(errors) > 0 {
		// Aggregate all errors and return
		return data, fmt.Errorf("one or more requests failed: %+v", errors)
	}

	// Collect data from dataChan
	for batch := range dataChan {
		data = append(data, batch...)
	}

	return data, nil
}
