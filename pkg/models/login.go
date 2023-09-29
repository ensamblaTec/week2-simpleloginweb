package models

import "github.com/ensamblaTec/learning/week/problema3/pkg/utils"

type Login struct {
	Email    string
	Password string
	Logged   bool
}

func CreateLogin(email, password string) *Login {
	hash := utils.Hash256FromString(password)
	return &Login{
		Email:    email,
		Password: hash,
	}
}
