package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func LinkVideosApi(r *gin.Engine) {
	videos := r.Group("/videos")

	videos.GET("/get-url/:filename", func(c *gin.Context) {
		filename := c.Param("filename")
		uri, err := GetPresignedVideoUrl(filename)
		if err != nil {
			GenerateBadResponse(c, http.StatusInternalServerError, err, "Hubo un problema al procesar su solicitud")
			return
		}
		GenerateGoodResponse(c, http.StatusOK, uri)
	})
}
