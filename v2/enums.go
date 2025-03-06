package mytools

import (
	"fmt"
	"reflect"
)

// MapEnum takes a map of interface_a to interface_b and a value of interface_a,
// then returns the corresponding interface_b value
func MapEnum[K comparable, V any](mapping map[K]V, value K) V {
	var defaultValue V
    targetType := reflect.TypeOf(defaultValue)

	strValue := fmt.Sprintf("%v", value)

	if mapping == nil {
		return reflect.ValueOf(strValue).Convert(targetType).Interface().(V)
	}

	if val, ok := mapping[value]; ok {
        return val
    }

	return reflect.ValueOf(strValue).Convert(targetType).Interface().(V)
}
