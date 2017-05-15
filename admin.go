package main

import (
	"net/http"
	"time"

	"github.com/labstack/echo"
)

func AdminLog(c echo.Context) error {
	return c.JSON(http.StatusOK, MapY(0, http.StatusText(200), time.Now().UTC()))
}
