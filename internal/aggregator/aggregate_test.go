package aggregator

import "testing"

type MockAggregator struct {
	Highs  []interface{}
	Lows   []interface{}
	Opens  []interface{}
	Closes []interface{}
}

func TestAggregate(t *testing.T) {
	tests := []struct {
		name   string
		input  []map[string]interface{}
		result []map[string]interface{}
	}{
		{
			"empty ",
			[]map[string]interface{}{},
			[]map[string]interface{}{},
		},
	}

	for _, entry := range tests {
		t.Run(entry.name, func(t *testing.T) {

		})
	}
}
