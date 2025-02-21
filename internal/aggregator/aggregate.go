package aggregator

import (
	"reflect"
	"stock-aggregator/utils"
)

func (a *Aggregator) Aggregate(data []map[string]interface{}) error {
	// Define a map to dynamically route keys to their respective slices
	keyGroups := map[string]*[]interface{}{
		"dayHigh": &a.Highs, "h": &a.Highs, "high": &a.Highs,
		"dayLow": &a.Lows, "l": &a.Lows, "low": &a.Lows,
		"open": &a.Opens, "o": &a.Opens,
		"previousClose": &a.Closes, "pc": &a.Closes, "close": &a.Closes,
	}

	// Iterate over data and append relevant values
	for _, item := range data {
		for key, value := range item {
			if slice, exists := keyGroups[key]; exists {
				*slice = append(*slice, value)
			}
		}
	}

	return nil
}

func (a *Aggregator) CalculateAverages() map[string]float64 {
	result := make(map[string]float64)

	// Get Aggregator struct
	v := reflect.ValueOf(a).Elem()

	// For each field in struct
	for i := 0; i < v.NumField(); i++ {
		// convert each field to interface
		field := v.Field(i).Interface()
		// get average of each fiels
		average := calculate(field)
		// reassign the field to the average value
		result[v.Type().Field(i).Name] = average
	}

	return result
}

func calculateSum(field interface{}) float64 {
	slice, ok := field.([]interface{})
	if !ok {
		return 0
	}

	var sum float64
	for _, val := range slice {
		if num, ok := val.(float64); ok {
			sum += num
		}
	}

	result := sum / float64(len(slice))

	return roundToTwoDecimals(result)
}

func roundToTwoDecimals(num float64) float64 {
	return math.Round(num*100) / 100
}
