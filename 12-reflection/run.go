package main

import (
	"reflect"
)

func run(x interface{}, fn func(field string)) {
	value := getValue(x)

	runValue := func(value reflect.Value) {
		run(value.Interface(), fn)
	}

	switch value.Kind() {
	case reflect.String:
		fn(value.String())
	case reflect.Struct:
		for i := 0; i < value.NumField(); i++ {
			runValue(value.Field(i))
		}
	case reflect.Slice, reflect.Array:
		for i := 0; i < value.Len(); i++ {
			runValue(value.Index(i))
		}
	case reflect.Map:
		for _, key := range value.MapKeys() {
			run(value.MapIndex(key).Interface(), fn)
		}
	}
}

func getValue(x interface{}) reflect.Value {
	value := reflect.ValueOf(x)

	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}

	return value
}
