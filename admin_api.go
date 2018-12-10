package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func LinkAdminAPI(r *gin.Engine) {
	admin := r.Group("/admin")
	bAuth := gin.BasicAuth(
		gin.Accounts{
			"admin": "cita2k18",
		},
	)
	admin.Use(bAuth)

	admin.GET("/home", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello CITA Admin")
	})

	admin.GET("/all-videos", func(c *gin.Context) {
		videos, err := GetAllVideos()
		if err != nil {
			GenerateBadResponse(c, http.StatusInternalServerError, err, "Hubo un problema con nuestros servidores")
			return
		}

		type MiniVideo struct {
			Title string `json:"title"`
			Longitude string `json:"longitude"`
			Latitude string `json:"latitude"`
			Creator string `json:"creator"`
		}

		finalVideos := make([]*MiniVideo, 0)
		for _, v := range videos {
			finalVideos = append(finalVideos,
				&MiniVideo{
					Title:v.Title,
					Creator:strings.Replace(strings.Replace(v.OwnerID.String(), "ObjectIdHex(\"", "", -1), "\")", "", -1),
					Latitude: v.Latitude,
					Longitude: v.Longitude,
				},
			)
		}

		GenerateGoodResponse(c, http.StatusOK, finalVideos)
	})
}