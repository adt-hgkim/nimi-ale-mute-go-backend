package service

import "golang.org/x/crypto/bcrypt"

type hashStruct struct {
	cost int
}

func (h hashStruct) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), h.cost)
	return string(bytes), err
}

func (hashStruct) CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

var Hash = hashStruct{cost: 14}
