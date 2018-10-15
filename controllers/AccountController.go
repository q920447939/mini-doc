package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"wahaha/utils"
	"fmt"
)

//注册页面
func RegisteredHtml(context *gin.Context)  {
	baseName, _ := utils.AppConfigMap[utils.BASE_NAME]
	htmlTitle, _ := utils.AppConfigMap[utils.HTML_TITLE]
	baseUrl, _ := utils.AppConfigMap[utils.BASE_URL]
	context.HTML(http.StatusOK, "register.html", gin.H{
		"baseName":            baseName,
		"registeredHTMLTitle": htmlTitle,
		"baseUrl": baseUrl,
	})

}


//注册
func Add(context *gin.Context)  {
	type0 := context.DefaultPostForm("type", "alert")//可设置默认值
	 msg := context.PostForm("password1")
	fmt.Printf("type:%v,msg:%v",type0,msg)

}
