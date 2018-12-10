package main

import (
	"errors"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"net/http"
)

func LinkOpenAPI(r *gin.Engine) {
	open := r.Group("/open")
	open.POST("/:id/forgot-password", func(c *gin.Context) {
		id := c.Param("id")
		if !(govalidator.IsNumeric(id) && len(id)==8) {
			GenerateBadResponse(c, http.StatusBadRequest, errors.New("invalid id"), "DNI invalido")
			return
		}
		user, err := GetUserByDNI(id)
		if err != nil {
			GenerateBadResponse(c, http.StatusBadRequest, err, "Problemas con el servidor")
			return
		}

		if user.Password == "" {
			user.Password = user.PseudoPassword
		}
		r := NewRequest([]string{user.Email}, "Password Recovery")
		err = r.Send("./templates/forgotpassword.html", map[string]interface{}{
			"Name": user.Name,
			"Password": user.Password,
		})

		if err != nil {
			GenerateBadResponse(c, http.StatusBadRequest, err, "No se pudo enviar el mensaje")
			return
		}

		GenerateGoodResponse(c, http.StatusOK, "ok, email sent")
	})
}