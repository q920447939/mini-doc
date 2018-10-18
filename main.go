package main

import (

	"github.com/gin-gonic/gin"
	"net/http"
	"wahaha/filters/auth"
	mG "wahaha/init/gin"
	"wahaha/routers"
	_ "wahaha/connections/database/mysql"
	_ "wahaha/connections/redis"
)

func main() {
	router := gin.New()
	mG.InitGin(router)
	routers.GinRouter()
	router.Static("/static", "./static")
	//router.StaticFS("/more_static", http.Dir("my_file_system"))
	//router.StaticFile("/favicon.ico", "./resources/favicon.ico")

	router.Run()
}

func JwtSetExample(c *gin.Context) {
	authDr, _ := c.MustGet("jwt_auth").(*auth.Auth)

	token, _ := (*authDr).Login(c.Request, c.Writer, map[string]interface{}{"id": "123"}).(string)

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "ok",
		"data": gin.H{
			"token": token,
		},
	})
}
