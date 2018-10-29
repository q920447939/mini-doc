package routers

import (
	"wahaha/controllers"
	"wahaha/module/gin"
	ig "wahaha/module/gin"
	ggin "github.com/gin-gonic/gin"
	"net/http"
	)

func GinRouter() {

	gin.GinEngine.GET("/register", controllers.RegisteredHtml)

	gin.GinEngine.GET("/", func(context *ggin.Context) {
 		context.HTML(http.StatusOK, "index.html", ggin.H{
			"title": "Main website",
		})
	})

	userGroup := ig.GinEngine.Group("/user")
	{
		userGroup.POST("/add", controllers.Register)
		userGroup.POST("/login", controllers.Login)
	}


	menus := ig.GinEngine.Group("/menus")
	{
		viewMenus := menus.Group("/view")
		{
			viewMenus.GET("/top",controllers.FindIndexMenus)
		}
	}

}
