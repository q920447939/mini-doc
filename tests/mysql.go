package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"wahaha/module/rbac"
	"fmt"
)

func main() {
	//db, err := gorm.Open("mysql", "root:123456@/gin-template?charset=utf8&parseTime=True&loc=Local")
	db,err:=gorm.Open("mysql","root:123456@tcp(127.0.0.1:3306)/gin-template?charset=utf8")

	if err != nil {
		fmt.Println(err,"---")
	}
	db.AutoMigrate()
	db.CreateTable(&rbac.Member{})
}
