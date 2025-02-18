package fetcher

import (
	"fmt"
)

func (f *Fetcher) Fetch() ([]map[string]interface{}, error) {
	dataChan, errChan := f.owner()
	data, err := f.consumer(dataChan, errChan)

	if err != nil {
		return nil, err
	}

	return data, nil
}

func (f *Fetcher) owner() (<-chan []map[string]interface{}, <-chan error) {
	dataChan := make(chan []map[string]interface{}, len(f.Urls))
	errChan := make(chan error, len(f.Urls))

	go func() {
		for _, url := range f.Urls {

			go func(url string) {
				batch, err := f.RequesterParser.RequestAndParse(url)

				if err != nil {
					f.ErrorLog.Printf("\t\nurl: %s\t\nerr: %+v", url, err)
					errChan <- err
					return
				}

				dataChan <- batch
			}(url)
		}
	}()

	return dataChan, errChan
}

func (f *Fetcher) consumer(dataChan <-chan []map[string]interface{}, errChan <-chan error) ([]map[string]interface{}, error) {
	// Collect results from the channels
	var data []map[string]interface{}
	var errors []error

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
