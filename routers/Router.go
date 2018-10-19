package routers

import (
	"wahaha/controllers"
	"wahaha/module/gin"
	ig "wahaha/module/gin"
		ggin "github.com/gin-gonic/gin"
	"fmt"
)

func GinRouter() {

	gin.GinEngine.GET("/register", controllers.RegisteredHtml)
	userGroup := ig.GinEngine.Group("/user")
	{
		userGroup.POST("/add", controllers.Register)
		userGroup.POST("/login", controllers.Login)
	}

	gin.GinEngine.GET("/add", ada)
}

func ada(c *ggin.Context) {
	// cache存储
/*	cacheStore, _ := c.MustGet(cache.CACHE_MIDDLEWARE_KEY).(*persistence.CacheStore)
	(*cacheStore).Set("abcdefg", "abcdefgabcdefg", time.Minute)
//	(*cacheStore).Delete("key")

	c.JSON(http.StatusOK, ggin.H{
		"code": 0,
		"msg":  "ok",
		"data": ggin.H{
			"store_result": "success",
		},
	})*/
	/*authDr, _ := c.MustGet("web_auth").(*auth.Auth)
	(*authDr).Login(c.Request, c.Writer, map[string]interface{}{"id": "iiiiiiiiiiiiiiiiii"})
	ggin.
	// 返回html
	c.HTML(http.StatusOK, "index.tpl", ggin.H{
		"title": "login success!",
	})*/
	if  ss{
		c.SetCookie("name","test",30,"/","localhost",false,true)
		c.String(200, "Cookie:%s")
		ss = false
	}else {
		s, _ := c.Cookie("name")
		fmt.Println(s)
	}

}

var ss  = true
