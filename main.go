package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	mG "wahaha/init/gin"
	"wahaha/routers"
	_ "wahaha/connections/database/mysql"
	_ "wahaha/connections/redis"
	"wahaha/connections/redis"
	"strings"
	mmodule_gin "wahaha/module/gin"
	"time"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/ScottHuangZL/gin-jwt-session"
)

func main() {
	router := gin.New()

	router.Use(session.ClearMiddleware()) //important to avoid mem leak

	router.Static("/static", "./static")
	router.StaticFS("/more_static", http.Dir("my_file_system"))
	router.StaticFile("/favicon.ico", "./resources/favicon.ico")
	mG.InitGin(router)
	routers.GinRouter()


	store := sessions.NewCookieStore([]byte("secret"))
	store.Options(sessions.Options{
		MaxAge: int(30 * time.Minute), //30min
		Path:   "/",
	})
	router.Use(sessions.Sessions("aaaaaaaaaaaa", store))
	mmodule_gin.Run(router)

}

func ValidToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		/*if cookie, err := c.Request.Cookie("session_id"); err == nil {
			value := cookie.Value
			fmt.Println(value)
			if value == "onion" {
				c.Next()
				return
			}
		}
		if url := c.Request.URL.String(); url == "/add" {
			c.Next()
			return
		}
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
		c.Abort()*/
		url := c.Request.URL.String()
		if url == "/register" || url == "/user/add" {
			c.Next()
			return
		} else {
			for _, v := range FilterMap {
				siles := strings.Split(string(v), ".")
				if len(siles) > 1 {
					s := FilterMap[siles[len(siles)-1] ]
					if s != "" {
						c.Next()
					}
				}
			}
			checkToken(c)
		}
		return
	}

}

const (
	FilterCss  = ".css"
	FilterJs   = ".js"
	FilterHtml = ".html"
)

var FilterMap = map[string]string{
	FilterCss:  FilterCss,
	FilterJs:   FilterJs,
	FilterHtml: FilterHtml,
}

func checkToken(c *gin.Context) {
	access_token := c.Param("access_token")
	if access_token == "" {
		c.JSON(http.StatusOK, gin.H{
			"message": "access_token is not empty!",
		})
		return
	}
	Rvalue, _ := redis.Client.Get(access_token)
	if Rvalue == "" {
		c.JSON(http.StatusOK, gin.H{
			"message": "access_token is fake or expire!",
		})
		return
	}

}
