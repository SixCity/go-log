package main

import (
	"net/http"

	"fmt"

	"encoding/json"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func ServerHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderServer, "city/1.1")
		return next(c)
	}
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

		var data []Animal
		_ = json.Unmarshal(jsonBlob, &data)

		A3 := JSM(data)

		return c.JSON(http.StatusOK, St2Map(A3))
	})
	api.POST("/log", func(c echo.Context) error {

		req := &PersonInfo{
			Code:    1,
			Data:    "sss",
			Message: "eee",
		}

		req2, _ := json.Marshal(req)

		fmt.Println(req)
		fmt.Println(string(req2))

		map1 := make(map[string]interface{})
		map1["Code"] = "hello"
		map1["Data"] = "world"
		map1["Message"] = string(req2)

		str, err := json.Marshal(map1)

		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("map to json", string(str))

		//Celt()

		return c.JSON(http.StatusOK, map1)
	})

	e.Logger.Fatal(e.Start(":1222"))
}
