package routers

import (
			"wahaha/controllers"
	"wahaha/module/gin"
			ig "wahaha/module/gin"
)

func GinRouter() {

	gin.GinEngine.GET("/register", controllers.RegisteredHtml)
	userGroup := ig.GinEngine.Group("/user")
	{
		userGroup.POST("/add", controllers.Register)
	}

}
