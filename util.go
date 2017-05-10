package main

import "reflect"

type PersonInfo struct {
	Code    int8        `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

// JSM is
func JSM(data interface{}) (req PersonInfo) {

	req = PersonInfo{
		Code:    0,
		Data:    data,
		Message: "word",
	}

	return req
}

//St2Map is
func St2Map(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		data[t.Field(i).Name] = v.Field(i).Interface()
	}
	return data
}
