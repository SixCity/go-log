package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var Idb *gorm.DB

func init() {
	var err error

	Idb, err = gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/gotest?charset=utf8&parseTime=True")
	//Idb, err := gorm.Open("mysql", "root:root@/gotest?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	//Idb2,_ = gorm.Open("mysql", "root:root@/gotest?charset=utf8&parseTime=True&loc=Local")
	defer Idb.Close()
	// Migrate the schema
	Idb.AutoMigrate(&Product{})
	// Create
	Idb.Create(&Product{Code: "L1212", Price: 1000})

	// Read
	var product Product
	//Idb.First(&product, 1) // find product with id 1

	Idb.Find(&product)

	fmt.Println(product)
}
