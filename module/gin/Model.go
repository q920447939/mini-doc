package gin

import (
	"github.com/jinzhu/gorm"
	config "wahaha/conf"
	)

var Model *gorm.DB

func init() {
	var err error
	Model, err = gorm.Open("mysql", config.GetEnv().DatebaseConfig.DATABASE_USERNAME+
		":"+config.GetEnv().DatebaseConfig.DATABASE_PASSWORD+"@tcp("+config.GetEnv().DatebaseConfig.DATABASE_IP+
		":"+config.GetEnv().DatebaseConfig.DATABASE_PORT+")/"+config.GetEnv().DatebaseConfig.DATABASE_NAME+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	Model.SingularTable(true)
}
