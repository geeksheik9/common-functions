package utils

import (
	"reflect"
)

func GetFields(config interface{}) []string {
	fields := []string{}
	for i := 0; i < reflect.ValueOf(config).NumField(); i++ {
		field := reflect.ValueOf(config).Type().Field(i).Name
		fields = append(fields, field)
	}
	return fields
}

func ShrinkArray(array []interface{}) []interface{} {
	for i := 0; i < len(array); i++ {
		if array[i] == nil {
			array[i] = array[len(array)-1]
			array[len(array)-1] = nil
			array = array[:len(array)-1]
		}
	}
	return array
}
