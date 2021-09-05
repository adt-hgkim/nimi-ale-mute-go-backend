package plugin

import (
	"github.com/adt-hgkim/nimi-ale-mute/api"
	"github.com/gin-gonic/gin"
)

func InitRouter(router *gin.Engine) {
	// USERS
	router.POST("users", api.Users.Post)

	// USERS/:id
	//router.PATCH()
	//router.DELETE()

	// AUTHS
	router.GET("auth", api.Auth.Get)
}
