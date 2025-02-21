package fetcher

import (
	"stock-aggregator/testdata"
	"testing"
)

func TestRequestAndParse(t *testing.T) {
	tests := []struct {
		name          string
		url           string
		expected      []map[string]interface{}
		expectedError bool
	}{
		{
			name:          "Invalid URL",
			url:           "https://foo-bar.foo/api/bar?page=2",
			expected:      nil,
			expectedError: true,
		},
		{
			name:          "ReqResUsers list of users",
			url:           "https://reqres.in/api/users?page=2",
			expected:      testdata.ReqResUsers,
			expectedError: false,
		},
		{
			name:          "JSONPlaceholder list of users",
			url:           "https://jsonplaceholder.typicode.com/users",
			expected:      testdata.JsonPlaceholderUsers,
			expectedError: false,
		},
	}

	for _, entry := range tests {
		f := &Fetcher{
			RequesterParser: &Requester{},
		}

		// per each entry make a request to the url provided
		// compare data with the expected data
		result, err := f.RequesterParser.RequestAndParse(entry.url)
		if err != nil && !entry.expectedError {
			t.Errorf("expected no error, but received: %+v", err)
		}

		if err == nil && entry.expectedError {
			t.Error("expected an error, but did not receive any")
		}

		// Debug the result to compare in detail
		if len(result) != len(entry.expected) {
			t.Errorf("expected length of the result to be %d, received: %d", len(entry.expected), len(result))
		}
	}
}
