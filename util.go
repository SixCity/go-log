package main

import (
	"fmt"
	"reflect"
)

// MapY is ...
func MapY(code int, message string, data interface{}) (iMap map[string]interface{}) {
	iMap = make(map[string]interface{})

	iMap["code"] = code
	iMap["message"] = message
	iMap["data"] = data

	return iMap
}

// St2Map is ...
func St2Map(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		data[t.Field(i).Name] = v.Field(i).Interface()
	}
	return data
}

// CheckErr is ...
func CheckErr(err error) {

	if err != nil {

		fmt.Println(err)

		panic(err)

	}

}
