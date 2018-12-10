package main

import (
	"github.com/gin-gonic/gin"
	"github.com/appleboy/gin-jwt"
	"github.com/asaskevich/govalidator"
	"net/http"
	"github.com/pkg/errors"
	"io/ioutil"
	"log"

)

func LinkUsersHelper(r *gin.Engine) {
	users := r.Group("/users")
	users.GET("/:id", func(c *gin.Context) {
		id := c.Param("id")
		claims := jwt.ExtractClaims(c)

		if !(govalidator.IsNumeric(id) && len(id)==8) {
			GenerateBadResponse(c, http.StatusBadRequest, errors.New("invalid id"), "DNI invalido")
			return
		}

		if id != claims["id"] {
			GenerateBadResponse(c, http.StatusForbidden, errors.New("unauthorized credentials, please sign in"),
				"Usted no está autorizado para acceder a estos datos")
			return
		}

		user, err := GetUserByDNI(id)

		if err != nil {
			GenerateBadResponse(c, http.StatusInternalServerError, err, err.Error())
			return
		}

		GenerateGoodResponse(c, http.StatusOK, user)
	})

	users.GET("/:id/videos", func(c *gin.Context) {
		id := c.Param("id")
		claims := jwt.ExtractClaims(c)

		if !(govalidator.IsNumeric(id) && len(id)==8) {
			GenerateBadResponse(c, http.StatusBadRequest, errors.New("invalid id"), "DNI invalido")
			return
		}

		if id != claims["id"] {
			GenerateBadResponse(c, http.StatusForbidden, errors.New("unauthorized credentials, please sign in"),
				"Usted no está autorizado para acceder a estos datos")
			return
		}

		videos, err := GetListOfVideoOfUserByDni(id)

		if err != nil {
			GenerateBadResponse(c, http.StatusInternalServerError, err, err.Error())
			return
		}

		GenerateGoodResponse(c, http.StatusOK, videos)

	})

	users.POST("/:id/save-video", func(c *gin.Context) {
		id := c.Param("id")
		claims := jwt.ExtractClaims(c)

		if !(govalidator.IsNumeric(id) && len(id)==8) {
			GenerateBadResponse(c, http.StatusBadRequest, errors.New("invalid id"), "DNI invalido")
			return
		}

		if id != claims["id"] {
			GenerateBadResponse(c, http.StatusForbidden, errors.New("unauthorized credentials, please sign in"),
				"Usted no está autorizado para acceder a estos datos")
			return
		}

		file, err := c.FormFile("video")
		if err != nil {
			GenerateBadResponse(c, http.StatusInternalServerError, err, "Hubo un problema al recibir el video")
		    return
		}

		log.Println("bottle neck 1")
		video, err := file.Open()

		if err != nil {
			GenerateBadResponse(c, http.StatusInternalServerError, err, "Hubo un problema al abir el video")
			return
		}
		log.Println("bottle neck 2")
		data, err := ioutil.ReadAll(video)
		if err != nil {
			GenerateBadResponse(c, http.StatusInternalServerError, err, "Hubo un problema al leer el video")
			return
		}
		log.Println("bottle neck 3")


		thumb, err := GetThumbnailFromVideoBytes(data)
		if err != nil {
			GenerateBadResponse(c, http.StatusInternalServerError, err, "No se pudo generar el thumbnail del video")
		    return
		}

		d, err := ioutil.ReadAll(thumb)
		if err != nil {
			GenerateBadResponse(c, http.StatusInternalServerError, err, "No se pudo generar el thumbnail del video")
			return
		}

		title := c.PostForm("title")
		latitude := c.PostForm("latitude")
		longitude := c.PostForm("longitude")
		location := c.PostForm("location")
		url := file.Filename


		videoObject := GetVideoFromBaseData(title, latitude, longitude, location, url, d)

		err = CreateAndLinkVideoWithUserByDNI(videoObject, id)
		if err != nil {
			GenerateBadResponse(c, http.StatusInternalServerError, err, "No se pudo añadir video a la cuenta del usuario")
			return
		}

		log.Printf("Uploading %s with %d bytes", file.Filename, file.Size)
		err = SaveNewVideo(file.Filename, "video/mp4", data, c)
		if err != nil {
			GenerateBadResponse(c, http.StatusInternalServerError, err, "Hubo un problema al intentar guardar el video")
			return
		}

		log.Println("bottle neck 4")

		GenerateGoodResponse(c, http.StatusOK, "video uploaded", "Gracias por colaborar")
	})

}
