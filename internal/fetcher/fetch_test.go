package fetcher

import (
	"fmt"
	"log"
	"testing"
)

type MockRequesterParser struct {
	MockRequest func(url string) ([]map[string]interface{}, error)
}

func (mrp *MockRequesterParser) RequestAndParse(url string) ([]map[string]interface{}, error) {
	return mrp.MockRequest(url)
}

type MockFetcher struct {
	ErrorLog        *log.Logger
	Urls            *[]string
	RequesterParser MockRequesterParser
}

func TestFetch(t *testing.T) {
	tests := []struct {
		name        string
		testUrls    []string
		mockRequest func(url string) ([]map[string]interface{}, error)
		expectError bool
	}{
		{
			name:     "Successful fetch",
			testUrls: []string{"https://i-am-url.com", "https://i-am-stock-api.com", "https://go-dev.go"},
			mockRequest: func(url string) ([]map[string]interface{}, error) {
				return []map[string]interface{}{
					{"symbol": "AAPL"},
				}, nil
			},
			expectError: false,
		},
		{
			name:     "Unsuccessful fetch",
			testUrls: []string{"https://i-am-url.com", "https://i-am-stock-api.com", "https://go-dev.go"},
			mockRequest: func(url string) ([]map[string]interface{}, error) {
				return nil,
					fmt.Errorf("error occurred")
			},
			expectError: true,
		},
	}

	for _, entry := range tests {
		t.Run(entry.name, func(t *testing.T) {
			mockRequesterParser := &MockRequesterParser{
				MockRequest: entry.mockRequest,
			}

			fetcher := &Fetcher{
				ErrorLog:        log.Default(),
				Urls:            entry.testUrls,
				RequesterParser: mockRequesterParser,
			}

			data, err := fetcher.Fetch()
			if err != nil && entry.expectError == true {
				t.Errorf("expected error: %v, got: %v", entry.expectError, err)
			}

			if err == nil && len(data) == 0 {
				t.Errorf("expected data, got none")
			}
		})
	}
}
