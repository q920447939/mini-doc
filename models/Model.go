package models

import (
	"github.com/jinzhu/gorm"
	config "wahaha/conf"
	"fmt"
)

var Model *gorm.DB

func init() {
	var err error
	Model, err = gorm.Open("mysql", config.GetEnv().DATABASE_USERNAME+
		":"+config.GetEnv().DATABASE_PASSWORD+"@tcp("+config.GetEnv().DATABASE_IP+
		":"+config.GetEnv().DATABASE_PORT+")/"+config.GetEnv().DATABASE_NAME)
	fmt.Println("mysql", config.GetEnv().DATABASE_USERNAME+
		":"+config.GetEnv().DATABASE_PASSWORD+"@tcp("+config.GetEnv().DATABASE_IP+
		":"+config.GetEnv().DATABASE_PORT+")/"+config.GetEnv().DATABASE_NAME)
	if err != nil {
		panic(err)
	}
}
