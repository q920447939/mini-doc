package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
)

type TestController struct {
	beego.Controller
}

func (this *TestController) Test ()  {
	fmt.Println(4)
	this.Data["asdas"] = 11
	this.ServeJSON()
}