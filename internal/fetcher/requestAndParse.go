package fetcher

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Requester struct{}

func (r *Requester) RequestAndParse(url string) ([]map[string]interface{}, error) {
	body, err := request(url)
	if err != nil {
		return nil, err
	}
	parsedBody, err := parse(body)
	if err != nil {
		return nil, err
	}

	return parsedBody, nil
}

func request(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusMultipleChoices {
		return nil, fmt.Errorf("request failed with status: %d", resp.StatusCode)
	}

	// Read the response body into a byte slice
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	return body, nil
}

func parse(body []byte) ([]map[string]interface{}, error) {
	// Try to unmarshal the response as an array first
	var dataArray []map[string]interface{}
	if err := json.Unmarshal(body, &dataArray); err == nil {
		return dataArray, nil
	}

	// If unmarshaling as an array fails, try to unmarshal as an object
	var dataObject map[string]interface{}
	if err := json.Unmarshal(body, &dataObject); err == nil {
		// Check if the object contains a 'data' key:
		// Will handle MarketStack API response
		if data, exists := dataObject["data"]; exists {
			if dataArray, ok := data.([]interface{}); ok {
				convertedToMap, err := convertToMapSlice(dataArray)
				if err != nil {
					return nil, err
				}
				return convertedToMap, nil
			}

			return dataArray, nil
		}
		// If no 'data' key exists, treat the entire object as the result
		return []map[string]interface{}{dataObject}, nil
	}

	// If both attempts fail, return an error
	return nil, fmt.Errorf("failed to parse JSON. response is neither an array nor an object")
}
