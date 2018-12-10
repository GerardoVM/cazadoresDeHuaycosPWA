package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

func GenerateGoodResponse(c *gin.Context, code int, data interface{}, message... string) {
	msg := "ok"
	if len(message) >0 {
		msg = message[0]
	}
	c.JSON(code, Response{
		Data: data,
		Code: code,
		UserMessage: msg,
	})
}

func GenerateBadResponse(c *gin.Context, code int, err error, message string) {
	if gin.Mode() == gin.DebugMode {
		log.Println(err)
		log.Println(message)
	}

	c.JSON(code, Response{
		Error: err,
		Code: code,
		UserMessage: message,
	})
}