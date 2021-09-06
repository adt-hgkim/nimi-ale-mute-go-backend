package api

import (
	. "github.com/adt-hgkim/nimi-ale-mute/database"
	"github.com/adt-hgkim/nimi-ale-mute/service"
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
		service.Response(context, http.StatusBadRequest, err.Error(), nil)
		return

	}

	hashedPassword, _ := service.Hash.HashPassword(body.Password)

	user := User{Email: body.Email, Password: hashedPassword}
	response := DB.Create(&user)

	if response.Error != nil {
		service.Response(context, http.StatusUnauthorized, response.Error.Error(), nil)
		return
	}

	service.Response(context, http.StatusOK, "Signed up", nil)
	return
}

func (usersStruct) SoftDelete(context *gin.Context) {
	type URI struct {
		ID string `uri:"id" binding:"required"`
	}
	var uri URI

	if err := context.ShouldBindUri(&uri); err != nil {
		service.Response(context, http.StatusBadRequest, err.Error(), nil)
		return
	}

	var user User
	if response := DB.Where("id = ?", uri.ID).First(&user); response.Error != nil {
		service.Response(context, http.StatusNotFound, response.Error.Error(), nil)
		return
	}

	if response := DB.Delete(&user); response.Error != nil {
		service.Response(context, http.StatusInternalServerError, response.Error.Error(), nil)
		return
	}

	service.Response(context, http.StatusOK, "Soft deleted", nil)
	return
}

var Users usersStruct
