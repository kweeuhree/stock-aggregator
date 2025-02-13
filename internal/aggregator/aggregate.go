package aggregator

import (
	"encoding/json"
	"fmt"
)

func (a *Aggregator) Aggregate(data []map[string]interface{}) error {
	// Pretty print JSON
	prettyJSON, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		a.ErrorLog.Println("Failed to format JSON:", err)
	} else {
		fmt.Println(string(prettyJSON))
	}
	return nil
}
