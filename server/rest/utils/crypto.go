package utils

import (
	"github.com/ArkamFahry/uploadnexus/server/rest/exceptions"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword() (string, error) {
	const Op exceptions.Op = "crypto.HashPassword"

	passwordHash, err := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
	if err != nil {
		return "", exceptions.NewError(Op, "failed to hash password", err)
	}

	return string(passwordHash), nil
}

func ComparePassword(hashPassword, password string) (bool, error) {
	const Op exceptions.Op = "crypto.ComparePassword"

	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
	if err != nil {
		return false, exceptions.NewError(Op, "failed to compare password", err)
	}

	return true, nil
}
