package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

func main() {

	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowMethods:     []string{"PUT", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "content-type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowAllOrigins: true,
	}))
	LinkVideosApi(r)
	LinkAdminAPI(r)
	LinkOpenAPI(r)
	LinkAuthJWT(r) // Making the auth context, all above this will be restricted
	LinkUsersHelper(r)
	r.Run(":4700")
}
