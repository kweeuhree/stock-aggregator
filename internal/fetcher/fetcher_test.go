package fetcher

import (
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

func TestNew(t *testing.T) {
	testUrls := []string{"https://i-am-url.com", "https://i-am-stock-api.com", "https://go-dev.go"}
	tests := []struct {
		name            string
		errorLog        *log.Logger
		urls            []string
		requesterParser RequesterParser
		want            *Fetcher
	}{
		{
			name:            "Successful initialization",
			errorLog:        log.Default(),
			urls:            testUrls,
			requesterParser: &MockRequesterParser{},
			want: &Fetcher{
				ErrorLog:        log.Default(),
				Urls:            testUrls,
				RequesterParser: &MockRequesterParser{},
			},
		},
	}

	for _, entry := range tests {
		t.Run(entry.name, func(t *testing.T) {
			result := New(entry.requesterParser, entry.errorLog, entry.urls)
			if result != entry.want {
				t.Log("failed initializing the Fetcher correctly")
			}
		})
	}
}
