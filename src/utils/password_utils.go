package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(pass string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), 14)
	if err != nil {
		panic(err)
	}
	return string(bytes)
}

func CheckPassword(pass string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pass))
	return err == nil
}
