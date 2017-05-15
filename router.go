package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func NewRouter() *echo.Echo {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(NotFind)
	e.Use(ServerHeader)
	//e.Use(ServerHeader)
	e.Static("/", "public")
	// api 开放路由组
	api := e.Group("/api", MiddlewareTime)
	api.POST("/time", HandTime)
	api.POST("/log", HandLog)

	// admin 管理组
	admin := e.Group("/admin", MiddlewareKey)
	admin.GET("/logs", AdminLog)

	e.HTTPErrorHandler = httpErrorHandler

	//e.Any("/*", func(c echo.Context) (err error) {
	//	err = echo.ErrNotFound
	//	return
	//})

	return e
}
