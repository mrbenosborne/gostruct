package gostruct

import (
	"errors"
	"reflect"
	"strconv"
)

// ToInterface pass a struct to retrieve values in an interface.
// Returns interface.
func ToInterface(v interface{}) ([]interface{}, error) {
	list := []interface{}{}
	a := reflect.ValueOf(v)
	if !a.IsValid() {
		return list, errors.New("failed to get reflection value")
	}
	for i := 0; i < a.NumField(); i++ {
		val := a.Field(i)
		switch val.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			list = append(list, strconv.FormatInt(val.Int(), 10))
		case reflect.String:
			list = append(list, val.String())
		case reflect.Float32, reflect.Float64:
			list = append(list, val.Float())
		case reflect.Bool:
			list = append(list, val.Bool())
		}
	}
	return list, nil
}

// Count returns the number of fields in a struct.
//
// Returns 0 if an error has occurred.
func Count(v interface{}) (int, error) {
	a := reflect.ValueOf(v)
	if !a.IsValid() {
		return 0, errors.New("failed to get reflection value")
	}
	return a.NumField(), nil
}
