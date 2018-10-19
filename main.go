package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"wahaha/filters/auth"
	mG "wahaha/init/gin"
	"wahaha/routers"
	_ "wahaha/connections/database/mysql"
	_ "wahaha/connections/redis"
	"fmt"
	"wahaha/connections/redis"
	"strings"
)

func main() {
	router := gin.New()
	router.Static("/static", "./static")
	router.StaticFS("/more_static", http.Dir("my_file_system"))
	router.StaticFile("/favicon.ico", "./resources/favicon.ico")
	router.Use(ValidToken())
	mG.InitGin(router)
	routers.GinRouter()
	router.Run()
}

func aaa() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		fmt.Println(123)
	}
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
		if url == "/user/add" {
			c.Next()
			return
		} else {
			for _, v := range FilterMap {
				siles := strings.Split(string(v),".")
				if len(siles) >1 {
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
	}
	Rvalue, _ := redis.Client.Get(access_token)
	if Rvalue == "" {
		c.JSON(http.StatusOK, gin.H{
			"message": "access_token is fake or expire!",
		})
	}

}
