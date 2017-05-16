package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func AdminLog(c echo.Context) error {

	recordLog := new([]RecordLogs)

	Idb.Order("created_at desc").Find(&recordLog)
	//Idb.Limit(2).Offset(2).Find(&recordLog)

	return c.JSON(http.StatusOK, MapY(0, http.StatusText(200), recordLog))

}
