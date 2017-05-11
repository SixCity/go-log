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

// MapY is ...
func MapY(code int, message string, data interface{}) (iMap map[string]interface{}) {
	iMap = make(map[string]interface{})

	iMap["code"] = code
	iMap["msg"] = message
	iMap["data"] = data

	return iMap
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
