package jwt

import (
	"wahaha/filters/auth"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/dgrijalva/jwt-go"
	"time"
	"wahaha/utils/jwt/setting"
)

var jwtSecret = []byte(setting.App{}.JwtSecret)

type Claims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Query("access_token")
		if token == "" {
			c.Redirect(302, "/view/login?redirect_url="+c.Request.URL.String())
			c.Abort()
			return
		} else {
			_, err := ParseToken(token)
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					c.JSON(http.StatusOK, gin.H{
						"code":    "40010",
						"message": " access_token is expire",
					})
					break
				default:
					c.JSON(http.StatusOK, gin.H{
						"code":    "40000",
						"message": "invaild access_token",
					})
				}
				c.Abort()
			}
		}
		c.Next()
	}
}

func GenerateToken(username, password string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour)

	claims := Claims{
		username,
		password,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "mini_doc",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
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
	fmt.Println(token)

}

func JwtGetExample(c *gin.Context) {
	/*authDr, _ := c.MustGet("jwt_auth").(*auth.Auth)
	user := (*authDr).User(c)
	info := (*authDr).User(c).(map[interface{}]interface{})

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "ok",
		"data": gin.H{
			"id": info["id"],
		},
	})*/
}
