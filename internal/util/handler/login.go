package handler

import (
	"log"

	"github.com/gin-gonic/gin"
)

func LoginCheck() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		log.Println("LoginCheck")
	}
}
