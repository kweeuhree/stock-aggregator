package fetcher

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
)

func (f *Fetcher) Fetch() ([]map[string]interface{}, error) {
	f.InfoLog.Println("Attempting to fetch...")
	urls, err := f.getUrls()
	if err != nil {
		return nil, err
	}

	data, err := f.fetchData(urls)
	if err != nil {
		f.ErrorLog.Fatal(err)
	}

	// Pretty print JSON
	prettyJSON, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		f.ErrorLog.Println("Failed to format JSON:", err)
	} else {
		fmt.Println(string(prettyJSON))
	}
	return data, nil
}

func (f *Fetcher) getUrls() ([]string, error) {
	marketStackApikey := os.Getenv("MARKET_STACK_API_KEY")
	fmpCloudApiKey := os.Getenv("FMP_CLOUD_API_KEY")

	if marketStackApikey == "" || fmpCloudApiKey == "" {
		return nil, fmt.Errorf("empty api key")
	}

	stocks := "AAPL,MSFT"

	// marketStackUrl := fmt.Sprintf("https://api.marketstack.com/v1/intraday/latest?access_key=%s&symbols=%s", marketStackApikey, stocks)

	fmpCloudURl := fmt.Sprintf("https://fmpcloud.io/api/v3/quote/%s?apikey=%s", stocks, fmpCloudApiKey)

	urls := []string{fmpCloudURl}

	return urls, nil
}

func (f *Fetcher) fetchData(urls []string) ([]map[string]interface{}, error) {
	f.InfoLog.Println("Fetching data concurrently...")

	wg := sync.WaitGroup{}

	dataChan := make(chan []map[string]interface{}, len(urls))
	errChan := make(chan error, len(urls))

	for _, url := range urls {
		wg.Add(1)

		go func(url string) {
			defer wg.Done()
			batch, err := f.request(url)

			if err != nil {
				f.ErrorLog.Printf("failed fetching.\t\nurl: %s\t\nerr: %+v", url, err)
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

func (f *Fetcher) request(url string) ([]map[string]interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("request to %s failed with status: %d", url, resp.StatusCode)
	}

	var data []map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, fmt.Errorf("failed to parse JSON from %s: %w", url, err)
	}
	// Pretty print JSON
	prettyJSON, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		f.ErrorLog.Println("Failed to format JSON:", err)
	} else {
		fmt.Println(string(prettyJSON))
	}

	return data, nil
}
