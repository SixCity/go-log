package main

import (
	_ "github.com/go-sql-driver/mysql"
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
	api.GET("/time", HandTime)
	api.POST("/log", HandLog)

	//db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/gotest?parseTime=true")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer db.Close()
	//
	//_, err = db.Exec("CREATE TABLE IF NOT EXISTS gotest.hello(world varchar(50))")
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//nowTime := "sadas"
	//rs, err := db.Exec("INSERT INTO gotest.hello(world) VALUES (" + nowTime + ")")
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//rowCount, err := rs.RowsAffected()
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//log.Printf("inserted %d rows", rowCount)

	//rows, err := db.Query("SELECT world FROM gotest.hello")
	//if err != nil {
	//	log.Fatalln(err)
	//}
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

	e.Logger.Fatal(e.Start(":1222"))
}
