package mytools

import (
	"fmt"
	"reflect"
)

// MapEnum takes a map of interface_a to interface_b and a value of interface_a,
// then returns the corresponding interface_b value
func MapEnum(mapping map[interface{}]interface{}, value interface{}) interface{} {
	// Get target type
	var targetType reflect.Type
	for _, v := range mapping {
		targetType = reflect.TypeOf(v)
		break
	}

	strValue := fmt.Sprintf("%v", value)

	if mapping == nil {
		return reflect.ValueOf(strValue).Convert(targetType).Interface()
	}

	for k, v := range mapping {
		// Use reflect.DeepEqual to compare interfaces since direct comparison might not work
		if reflect.DeepEqual(k, value) {
			return v
		}
	}

	return reflect.ValueOf(strValue).Convert(targetType).Interface()
}
