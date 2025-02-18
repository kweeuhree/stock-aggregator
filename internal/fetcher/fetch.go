package fetcher

import (
	"fmt"
)

func (f *Fetcher) Fetch() ([]map[string]interface{}, error) {

	dataChan := make(chan []map[string]interface{}, len(f.Urls))
	errChan := make(chan error, len(f.Urls))

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

	close(dataChan)
	close(errChan)

	// Check if there were any errors
	if len(errors) > 0 {
		// Aggregate all errors and return
		return data, fmt.Errorf("one or more requests failed: %+v", errors)
	}

	return data, nil
}
