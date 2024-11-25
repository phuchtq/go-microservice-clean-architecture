package utils

import "golang.org/x/crypto/bcrypt"

func IsLoginPasswordMatched(password string, inputPass string) bool {
	return bcrypt.CompareHashAndPassword([]byte(password), []byte(inputPass)) == nil
}
