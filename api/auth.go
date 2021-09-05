package api

import (
	. "github.com/adt-hgkim/nimi-ale-mute/database"
	"github.com/adt-hgkim/nimi-ale-mute/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type _ struct{}

func (_) Get(context *gin.Context) {
	type Body struct {
		Email    string `form:"email" json:"email" xml:"email"`
		Password string `form:"password" json:"password" xml:"password"`
		Access   string `form:"access" json:"access" xml:"access"`
		Refresh  string `form:"refresh" json:"refresh" xml:"refresh"`
	}
	var body Body
	if err := context.ShouldBindJSON(&body); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"status":  400,
			"message": "Bad request.",
			"value":   err.Error(),
		})
		return
	}

	user := User{}
	response := DB.Where("email = ?", body.Email).First(&user)
	if response.Error != nil {
		context.JSON(400, gin.H{
			"status":  http.StatusUnauthorized,
			"message": response.Error.Error(),
		})
		return
	}

	if service.Hash.CheckPasswordHash(body.Password, user.Password) == false {
		context.JSON(400, gin.H{
			"status":  http.StatusUnauthorized,
			"message": "Wrong password",
		})
		return
	}

	context.JSON(200, gin.H{
		"status":  http.StatusOK,
		"message": "Signed in",
	})
	return
}

var Auth _
