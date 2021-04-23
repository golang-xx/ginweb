package utils

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func Password(plainpwd string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(plainpwd), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
	}
	encodePWD := string(hash)
	return encodePWD
}

func CheckPassword(plainpwd, cryptepwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(cryptepwd), []byte(plainpwd))
	return err == nil
}
