package helper

import (
	"golang.org/x/crypto/bcrypt"
)

func ToHashString(s string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(s), 10)
	return string(bytes)
}
