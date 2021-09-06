package service

import "github.com/gin-gonic/gin"

func Response(context *gin.Context, status int, message string, value interface{}) {
	context.JSON(status, gin.H{
		"status":  status,
		"message": message,
		"value":   value,
	})
}
