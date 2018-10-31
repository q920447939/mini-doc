package routers

import (
	"wahaha/controllers"
	"wahaha/module/gin"
	ig "wahaha/module/gin"
	ggin "github.com/gin-gonic/gin"
	"net/http"
	"wahaha/utils/jwt"

	"fmt"
)

func GinRouter() {



	index := gin.GinEngine.Group("/minidoc")
	index.Use(jwt.JWT())
	{
		index.GET("/", func(context *ggin.Context) {
			context.HTML(http.StatusOK, "index.html", ggin.H{
			})
		})
	}

	userGroup := ig.GinEngine.Group("/user")
	{
		userGroup.POST("/add", controllers.Register)
		userGroup.POST("/login", controllers.Login)
	}

	menus := ig.GinEngine.Group("/menus")
	{
		viewMenus := menus.Group("/view")
		{
			viewMenus.GET("/top", controllers.FindIndexMenus)

		}
	}

	login := ig.GinEngine.Group("/login")
	{
		login.GET("/captcha",controllers.Captcha)
	}

	RouterHtml()

	book := ig.GinEngine.Group("/book")
	{
		book.GET("/list", controllers.List)
	}

	ig.GinEngine.GET("/jwts", JwtSet)


}

func RouterHtml() {
	htmlUseAuthView := ig.GinEngine.Group("/view")
	{
		htmlUseAuthView.GET("/login", func(context *ggin.Context) {
			context.HTML(http.StatusOK, "login.html", ggin.H{

			})
		})
		htmlUseAuthView.GET("/register",controllers.RegisteredHtml)
	}

}

func JwtSet(context *ggin.Context) {
	username := context.Query("username")
	password := context.Query("password")
	if username == "" {
		username = "username"
		password = "password"
	}
	if token, e := jwt.GenerateToken(username, password); e != nil {
		panic(e)
	} else {
		claims, i := jwt.ParseToken(token)
		fmt.Printf("claims = %v ,i = %v", claims, i)
		context.JSON(http.StatusGone, ggin.H{
			"data": token,
		})
	}
}
