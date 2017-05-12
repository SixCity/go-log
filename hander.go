package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

/*
time
*/
func HandTime(c echo.Context) error {
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

	var n [10]int /* n 是一个长度为 10 的数组 */
	var i, j int

	/* 为数组 n 初始化元素 */
	for i = 0; i < 10; i++ {
		n[i] = i + 100 /* 设置元素为 i + 100 */
	}

	/* 输出每个数组元素的值 */
	for j = 0; j < 10; j++ {
		fmt.Printf("Element[%d] = %d\n", j, n[j])
	}

	var array5 = [...]string{3: "Tom", 2: "Alice"}
	fmt.Printf("array5--- type:%T \n", array5)

	return c.JSON(http.StatusOK, MapY(http.StatusTooManyRequests, http.StatusText(401), data))
}

/*
日志
*/
func HandLog(c echo.Context) error {

	//rows, err := Idb.Query("SELECT * FROM user limit 10")
	//
	//defer rows.Close()
	//
	//CheckErr(err)
	//
	//columns, _ := rows.Columns()
	//
	//scanArgs := make([]interface{}, len(columns))
	//
	//values := make([]interface{}, len(columns))
	//
	//for j := range values {
	//
	//	scanArgs[j] = &values[j]
	//
	//}
	//
	//record := make(map[string]string)
	//
	//for rows.Next() {
	//
	//	//将行数据保存到record字典
	//	err = rows.Scan(scanArgs...)
	//
	//	for i, col := range values {
	//
	//		if col != nil {
	//			record[columns[i]] = string(col.([]byte))
	//
	//		}
	//	}
	//}

	//rows, err := Idb.Query("SELECT world FROM gotest.hello")
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//defer rows.Close()
	//
	//for rows.Next() {
	//	var s string
	//	err = rows.Scan(&s)
	//	if err != nil {
	//		log.Fatalln(err)
	//	}
	//	log.Printf("found row containing %q", s)
	//}
	//rows.Close()
	//if err = rows.Err(); err != nil {
	//	log.Fatal(err)
	//}

	//fmt.Println(record)

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
}
