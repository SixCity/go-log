package main

import (
	"net/http"
	"strconv"

	"fmt"

	"encoding/json"

	"reflect"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func ServerHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderServer, "city/1.1")
		return next(c)
	}
}

type personInfo struct {
	Code    int8        `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

//Struct2Map is
func Struct2Map(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		data[t.Field(i).Name] = v.Field(i).Interface()
	}
	return data
}

type AlipayRemoteReqStruct struct {
	Ono         string   `json:ono`
	OrderItem   []Item   `json:item`
	OrderRefund []Refund `json:refund`
}
type Item struct {
	Ono string `json:ono`
	Oid int    `json:oid`
}
type Refund struct {
	Ono     string `json:ono`
	Item    int    `json:item`
	Content string `json:content`
	Imgs    string `json:imgs`
	Status  string `json:status`
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	//e.Use(ServerHeader)
	e.Static("/", "public")
	// api 开放路由组
	api := e.Group("/api", ServerHeader)
	api.GET("/time", func(c echo.Context) error {
		//return c.String(http.StatusOK, "time")

		var jsonBlob = []byte(`[
			{ "Name" : "Platypus" , "Order" : "Monotremata" } ,
			{ "Name" : "Quoll" ,     "Order" : "Dasyuromorphia" }
			] `)
		type Animal struct {
			Name  string `json:"name"`
			Order string `json:"order"`
		}
		type A2 struct {
			Code int         `json:"code"`
			Data interface{} `json:"data"`
		}

		var animals []Animal
		err := json.Unmarshal(jsonBlob, &animals)

		A3 := A2{
			Code: 0,
			Data: animals,
		}
		fmt.Println(Struct2Map(A3))
		if err != nil {
			fmt.Println("error:", err)
		}
		fmt.Printf("%+v", animals)

		return c.JSON(http.StatusOK, Struct2Map(A3))
	})
	api.POST("/log", func(c echo.Context) error {
		//reqMap := map[personInfo]string{}
		//req := map[string]string{}
		req := &personInfo{
			Code:    1,
			Data:    "sss",
			Message: "eee",
		}

		req2, _ := json.Marshal(req)
		//fmt.Printf(string(req))
		fmt.Println(req)
		fmt.Println(string(req2))

		//req2 := Struct2Map(req)

		map1 := make(map[string]interface{})
		map1["Code"] = "hello"
		map1["Data"] = "world"
		map1["Message"] = string(req2)
		//map1["4"] = req2
		//return []byte
		str, err := json.Marshal(map1)

		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("map to json", string(str))

		var m AlipayRemoteReqStruct
		m.Ono = "12345"
		m.OrderItem = append(m.OrderItem, Item{Ono: "Shanghai_VPN", Oid: 1})
		m.OrderItem = append(m.OrderItem, Item{Ono: "Beijing_VPN", Oid: 2})
		for i := 1; i < 6; i++ {
			str := []byte("物品")
			str = strconv.AppendInt(str, int64(i), 10)
			orders := Item{Ono: string(str), Oid: i}
			m.OrderItem = append(m.OrderItem, orders)
		}
		bytes, _ := json.Marshal(m)
		fmt.Printf("json:m,%s\n", bytes)

		var js AlipayRemoteReqStruct
		err = json.Unmarshal(bytes, &js)

		fmt.Printf("json:type,%s\n", reflect.TypeOf(bytes))

		if err != nil {
			fmt.Printf("format err:%s\n", err.Error())
			return nil
		}

		return c.JSON(http.StatusOK, map1)
	})

	e.Logger.Fatal(e.Start(":1222"))
}
