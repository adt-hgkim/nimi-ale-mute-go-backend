package plugin

import (
	"github.com/adt-hgkim/nimi-ale-mute-monsi/api"
	"github.com/gin-gonic/gin"
)

func InitRouter(router *gin.Engine) {

	router.GET("", func(context *gin.Context) {

		context.JSON(200, gin.H{
			"message": "nimi ale mute monsi",
		})
	})

	// USERS
	router.POST("users", api.Users.Post)

	// USERS/:id
	//router.PATCH()
	//router.DELETE()

	// AUTHS
	router.GET("auth", api.Auth.Get)
}
