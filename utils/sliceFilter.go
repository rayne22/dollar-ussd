package utils

import "reflect"

// Filter is a dynamic filtering function that filters a slice of structs based on the specified field and condition.
func Filter(slice interface{}, field string, condition func(value interface{}) bool, structType reflect.Type) interface{} {
	sliceValue := reflect.ValueOf(slice)
	filtered := reflect.MakeSlice(reflect.SliceOf(structType), 0, 0)

	for i := 0; i < sliceValue.Len(); i++ {
		structValue := sliceValue.Index(i)
		fieldValue := structValue.FieldByName(field).Interface()

		if condition(fieldValue) {
			filtered = reflect.Append(filtered, structValue)
		}
	}

	return filtered.Interface()
}
