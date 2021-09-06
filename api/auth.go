package api

import (
	. "github.com/adt-hgkim/nimi-ale-mute/database"
	"github.com/adt-hgkim/nimi-ale-mute/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type authStruct struct{}

func (authStruct) Get(context *gin.Context) {
	type Body struct {
		Email    string `form:"email" json:"email" xml:"email"`
		Password string `form:"password" json:"password" xml:"password"`
		Access   string `form:"access" json:"access" xml:"access"`
		Refresh  string `form:"refresh" json:"refresh" xml:"refresh"`
	}
	var body Body
	if err := context.ShouldBindJSON(&body); err != nil {
		service.Response(context, http.StatusBadRequest, err.Error(), nil)
		return
	}

	user := User{}
	response := DB.Where("email = ?", body.Email).First(&user)
	if response.Error != nil {
		service.Response(context, http.StatusUnauthorized, response.Error.Error(), nil)
		return
	}

	if service.Hash.CheckPasswordHash(body.Password, user.Password) == false {
		service.Response(context, http.StatusUnauthorized, "Wrong password", nil)
		return
	}

	service.Response(context, http.StatusOK, "Signed in", nil)
	return
}

var Auth authStruct
