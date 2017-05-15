package main

import (
	"fmt"
	"net/http"

	"time"

	"github.com/labstack/echo"
)

/*
time
*/
func HandTime(c echo.Context) error {
	return c.JSON(http.StatusOK, MapY(0, http.StatusText(200), time.Now().UTC()))
}

/*
日志
*/
func HandLog(c echo.Context) error {

	//var recordLog RecordLogs
	recordLog := new(RecordLogs)

	err := c.Bind(recordLog)

	if err != nil {
		fmt.Println("error")
		fmt.Println(err)

	}

	fmt.Println(recordLog)

	if recordLog.Type == "" && recordLog.Content == "" {

		return c.JSON(422, MapY(422, "error", "Fields are empty"))
	}

	Idb.Create(&recordLog)

	return c.JSON(http.StatusOK, MapY(200, "ok", recordLog))

}
