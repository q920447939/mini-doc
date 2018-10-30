package gin

import (
	"github.com/jinzhu/gorm"
	config "wahaha/conf"
	)

var Model *gorm.DB

func init() {
	var err error
	Model, err = gorm.Open("mysql", config.GetEnv().DATABASE_USERNAME+
		":"+config.GetEnv().DATABASE_PASSWORD+"@tcp("+config.GetEnv().DATABASE_IP+
		":"+config.GetEnv().DATABASE_PORT+")/"+config.GetEnv().DATABASE_NAME+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	Model.SingularTable(true)
}
