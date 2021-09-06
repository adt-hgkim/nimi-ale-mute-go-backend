package plugin

import (
	"github.com/adt-hgkim/nimi-ale-mute/api"
	"github.com/gin-gonic/gin"
)

func InitRouter(router *gin.Engine) {
	router.POST("users", api.Users.Post)
	router.DELETE("users/:id", api.Users.SoftDelete)
	router.GET("auth", api.Auth.Get)
}
