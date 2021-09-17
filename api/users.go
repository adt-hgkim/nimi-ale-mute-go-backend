package api

import (
	. "github.com/adt-hgkim/nimi-ale-mute-monsi/database"
	"github.com/adt-hgkim/nimi-ale-mute-monsi/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type usersStruct struct{}

func (usersStruct) Post(context *gin.Context) {
	type Body struct {
		Email    string `form:"email" json:"email" xml:"email"  binding:"required"`
		Password string `form:"password" json:"password" xml:"password" binding:"required"`
	}
	var body Body
	if err := context.ShouldBindJSON(&body); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"status":  400,
			"message": err.Error(),
		})
		return
	}

	hashedPassword, _ := service.Hash.HashPassword(body.Password)

	user := User{Email: body.Email, Password: hashedPassword}
	response := DB.Create(&user)

	if response.Error != nil {
		context.JSON(400, gin.H{
			"status":  http.StatusUnauthorized,
			"message": response.Error.Error(),
		})
		return
	}

	context.JSON(200, gin.H{
		"status":  http.StatusOK,
		"message": "Signed up",
	})
	return
}

var Users usersStruct
