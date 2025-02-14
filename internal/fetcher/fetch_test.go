package fetcher

import (
	"fmt"
	"log"
	"testing"
)

func TestFetch(t *testing.T) {
	tests := []struct {
		name        string
		mockRequest func(url string) ([]map[string]interface{}, error)
		expectError bool
	}{
		{
			name: "Successful fetch",
			mockRequest: func(url string) ([]map[string]interface{}, error) {
				return []map[string]interface{}{
					{"symbol": "AAPL"},
				}, nil
			},
			expectError: false,
		},
		{
			name: "Unsuccessful fetch",
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
				Urls:            []string{"https://i-am-url.com", "https://i-am-stock-api.com", "https://go-dev.go"},
				RequesterParser: mockRequesterParser,
			}

			data, err := fetcher.Fetch()
			if err != nil && entry.expectError == false {
				t.Errorf("did not expect error, got: %v", err)
			}

			if err == nil && entry.expectError == true {
				t.Error("expected error, but did not get any")
			}

			if err == nil && len(data) == 0 {
				t.Errorf("expected data, got none")
			}
		})
	}
}
