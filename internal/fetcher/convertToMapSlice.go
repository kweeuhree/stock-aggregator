package fetcher

import "fmt"

func convertToMapSlice(data []interface{}) ([]map[string]interface{}, error) {

	var result []map[string]interface{}
	for _, item := range data {
		if m, ok := item.(map[string]interface{}); ok {
			result = append(result, m)
		} else {
			return nil, fmt.Errorf("unexpected type in 'data' array: expected map[string]interface{}")
		}
	}
	return result, nil
}
