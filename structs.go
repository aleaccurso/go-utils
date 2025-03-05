package mytools

import (
	"reflect"
)

// MapEnum takes a map of interface_a to interface_b and a value of interface_a,
// then returns the corresponding interface_b value
func MapEnum(mapping map[interface{}]interface{}, value interface{}) interface{} {
	// Check if mapping is nil
	if mapping == nil {
		return nil
	}

	// Check if the value exists in the map
	for k, v := range mapping {
		// Use reflect.DeepEqual to compare interfaces since direct comparison might not work
		if reflect.DeepEqual(k, value) {
			return v
		}
	}

	return nil
}
