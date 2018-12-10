package main

import (
	"github.com/gin-gonic/gin"
	"time"
	"github.com/appleboy/gin-jwt"
	jwt2 "gopkg.in/dgrijalva/jwt-go.v3"
	"net/http"
)

func authenticator(userDNI string, password string, c *gin.Context) (interface{}, bool) {
	err := CheckCredentials(userDNI, password)
	if err != nil {
		return err, false
	}

	user, err := GetUserByDNI(userDNI)
	if err != nil {
		return err, false
	}

	return user, true
}

func LinkAuthJWT(r *gin.Engine) {
	authMiddleware := &jwt.GinJWTMiddleware{
		Realm:      "citautec",
		Key:        []byte("cita"),
		Timeout:    30*24*time.Hour,
		MaxRefresh: 30*24*time.Hour,
		SendCookie: true,
		IdentityHandler: func(claims jwt2.MapClaims) interface{} {
			return jwt.MapClaims{
				"id": claims["id"],
				"exp": claims["exp"],
				"name": "Mr. Strange",
			}
		},
		Authenticator: authenticator,
		Authorizator: func(user interface{}, c *gin.Context) bool {
			return true
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		TokenLookup: "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		TimeFunc: time.Now,
	}

	r.POST("/signup", func(c *gin.Context) {

		data := new(NewUser)
		err := c.BindJSON(data)
		if err != nil {
			c.JSON(http.StatusInternalServerError, Response{
				Error: err,
				Code: 500,
				UserMessage: "Hubo un problema con los datos, por favor verifiquelos",
			})
			return
		}

		user, err := SignUpUser(data)
		if err != nil {
			c.JSON(http.StatusInternalServerError, Response{
				Error: err,
				Code: 500,
				UserMessage: "Datos incorrectos, por favor verifiquelos",
			})
			return
		}

		//authMiddleware.LoginHandler(c)

		//c.Header("Content-Type", "application/json")

		//c.Redirect(http.StatusOK, "/auth/login")

		c.JSON(http.StatusOK, Response{
			Code: 200,
			UserMessage: "Su usuario ha sido creado con éxito",
			Data: user,
		})

	})

	r.POST("/login", authMiddleware.LoginHandler)
	r.OPTIONS("/login", authMiddleware.LoginHandler)

	r.Use(authMiddleware.MiddlewareFunc())

	auth := r.Group("/auth")
	auth.GET("/hello", func (c *gin.Context) {
		claims := jwt.ExtractClaims(c)
		c.JSON(200, gin.H{
			"userID": claims["id"],
			"text":   "Hello World.",
		})
	})
	auth.GET("/refresh_token", authMiddleware.RefreshHandler)

}
