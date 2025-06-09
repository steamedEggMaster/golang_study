package service

import "golang.org/x/crypto/bcrypt"

func EncryptPassword(password string) (string, error) {
	encrypted, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(encrypted), err
}

func ComparePassword(encryptPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(encryptPassword), []byte(password))
	return err == nil
}
