package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/configor"
	"github.com/jinzhu/gorm"
)

var Idb *gorm.DB

var Config = struct {
	APPName string `default:"app name"`

	DB struct {
		Name     string
		User     string `default:"root"`
		Password string `required:"true" env:"DBPassword"`
		Port     string `default:"3306"`
		Host     string `default:"127.0.0.1"`
		Option   string
	}
}{}

func init() {
	var err error

	configor.Load(&Config, "config.yml")

	fmt.Printf("config : %#v", Config)

	uDB := Config.DB

	uri := uDB.User + ":" + uDB.Password + "@tcp(" + uDB.Host + ":" + uDB.Port + ")/" + uDB.Name + "?" + uDB.Option

	Idb, err = gorm.Open("mysql", uri)
	if err != nil {
		panic("failed to connect database")
	}
	defer Idb.Close()

}
