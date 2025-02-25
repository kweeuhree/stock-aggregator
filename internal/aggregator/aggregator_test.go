package aggregator

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name string
		size int
	}{
		{"One", 1},
		{"Long number", 1155},
		{"Very long number", 125555},
	}
	for _, entry := range tests {
		agg := New(entry.size)
		// Get Aggregator struct
		v := reflect.ValueOf(agg).Elem()
		// For each field in struct
		for i := 0; i < v.NumField(); i++ {
			// convert each field to interface
			fieldSize := v.Field(i).Cap()
			if fieldSize != entry.size {
				t.Errorf("expected Aggregator to initialize arrays with size %d, but received %d", entry.size, fieldSize)
			}
		}

	}
}
