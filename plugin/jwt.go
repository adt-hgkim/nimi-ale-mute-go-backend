package plugin

import (
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

type jwtStruct struct{}

func (jwtStruct) CreateAccessToken(userId uint32) (string, error) {
	if err := os.Setenv("ACCESS_SECRET", "tan-jan-aku"); err != nil {
		return "", err
	}
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = userId
	atClaims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return "", err
	}
	return token, nil
}

func (jwtStruct) CreateRefreshToken(userId uint32) (string, error) {
	if err := os.Setenv("REFRESH_SECRET", "tan-jan-aku"); err != nil {
		return "", err
	}
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = userId
	atClaims["exp"] = time.Now().Add(time.Hour * 24 * 7).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("REFRESH_SECRET")))
	if err != nil {
		return "", err
	}
	return token, nil
}

func (jwtStruct) CheckToken(userId uint32) (string, error) {
	return "", nil
}

var Jwt jwtStruct
