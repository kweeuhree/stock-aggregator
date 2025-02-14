package fetcher

import "testing"

func TestConvertToMapSlice(t *testing.T) {
	tests := []struct {
		name          string
		testData      []interface{}
		result        []map[string]interface{}
		expectedError bool
	}{
		{
			name: "Valid input",
			testData: []interface{}{
				map[string]interface{}{"hello": "world"},
				map[string]interface{}{"foo": "bar"},
			},
			result: []map[string]interface{}{
				{"hello": "world"},
				{"foo": "bar"},
			},
			expectedError: false,
		},
		{
			name: "Invalid input - not a map",
			testData: []interface{}{
				"not a map",
				map[string]interface{}{"foo": "bar"},
			},
			result:        nil,
			expectedError: true,
		},
	}

	for _, entry := range tests {
		t.Run(entry.name, func(t *testing.T) {
			result, err := convertToMapSlice(entry.testData)

			if len(result) != len(entry.result) {
				t.Errorf("expected result length: %d, got: %d", len(entry.result), len(result))
			}

			if err != nil && entry.expectedError == false {
				t.Errorf("expected no error, but received: %+v", err)
			}
		})
	}
}
