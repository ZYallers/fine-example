package handler

import (
	"log"

	"github.com/gin-gonic/gin"
)

func SignCheck() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		log.Println("SignCheck")
	}
}
