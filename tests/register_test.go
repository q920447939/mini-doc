package main

import (
	"testing"
	"fmt"
	"github.com/jinzhu/gorm"
	"wahaha/module/rbac"
	_ "wahaha/connections/database/mysql"

)

func Test_Register(t *testing.T)  {
	db, err := gorm.Open("mysql"," root:123456@tcp(127.0.0.1:3306)/gin-template")
	if err != nil {
		fmt.Println(err,"---")
	}
	db.CreateTable(&rbac.Member{})
}

func Test_RegisterB(t *testing.T)  {
	fmt.Println(14)

}