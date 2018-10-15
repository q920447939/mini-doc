package routers

import (
			"wahaha/controllers"
	"wahaha/models/gin"
			ig "wahaha/models/gin"
)

func GinRouter() {

	gin.GinEngine.GET("/register", controllers.RegisteredHtml)
	userGroup := ig.GinEngine.Group("/user")
	{
		userGroup.POST("/add", controllers.Add)
	}

}
