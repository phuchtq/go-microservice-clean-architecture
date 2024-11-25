package helper

import (
	envvar "architecture_template/constants/envVar"
	"architecture_template/constants/notis"
	"errors"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func ExtractDataFromToken(tokenString string, logger *log.Logger) (string, string, time.Time, error) {
	var errMsg error = errors.New(notis.GenericsErrorWarnMsg)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Unexpected signing method")
		}

		return []byte(os.Getenv(envvar.SecretKey)), nil
	})

	if err != nil {
		logger.Print("Error at ExtractDataFromToken - ", err)
		return "", "", GetPrimitiveTime(), errMsg
	}

	var userId string = ""
	var role string = ""
	var exp time.Time

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok || !token.Valid {
		return "", "", GetPrimitiveTime(), errMsg
	}

	if rawRole, ok := claims["role"].(string); rawRole == "" || !ok {
		return "", "", GetPrimitiveTime(), errMsg
	} else {
		role = rawRole
	}

	if id, ok := claims["userId"].(string); id == "" || !ok {
		return "", "", GetPrimitiveTime(), errMsg
	} else {
		userId = id
	}

	if expPeriod, ok := claims["exp"].(time.Time); expPeriod == GetPrimitiveTime() || !ok {
		return "", "", GetPrimitiveTime(), errMsg
	} else {
		exp = expPeriod
	}

	return userId, role, exp, nil
}
