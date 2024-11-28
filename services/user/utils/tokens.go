package utils

import (
	envvar "architecture_template/constants/envVar"
	"architecture_template/constants/notis"
	"architecture_template/helper"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateTokens(email string, userId string, role string, logger *log.Logger) (string, string, error) {
	var bytes = []byte(os.Getenv(envvar.SecretKey))
	var errMsg string = "Error while generating tokens - "

	accessToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": userId,
		"email":  email,
		"role":   role,
		"expire": time.Now().Add(NormalActionDuration).Unix(),
	}).SignedString(bytes)
	if err != nil {
		logger.Print(errMsg + fmt.Sprint(err))
		return "", "", errors.New(notis.InternalErr)
	}

	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": userId,
		"email":  email,
		"role":   role,
		"expire": time.Now().Add(RefreshTokenDuration).Unix(),
	}).SignedString(bytes)
	if err != nil {
		logger.Print(errMsg + fmt.Sprint(err))
		return "", "", errors.New(notis.InternalErr)
	}

	return accessToken, refreshToken, nil
}

func ExtractDataFromToken(tokenString string, logger *log.Logger) (string, string, time.Time, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Unexpected signing method")
		}

		return []byte(os.Getenv(envvar.SecretKey)), nil
	})

	if err != nil {
		logger.Println("Error at ExtractDataFromToken - ", err)
		return "", "", helper.GetPrimitiveTime(), errors.New(notis.GenericsErrorWarnMsg)
	}

	var userId string = ""
	var role string = ""
	var exp time.Time

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok || !token.Valid {
		return "", "", helper.GetPrimitiveTime(), errors.New(notis.GenericsErrorWarnMsg)
	}

	if rawRole, ok := claims["role"].(string); rawRole == "" || !ok {
		return "", "", helper.GetPrimitiveTime(), errors.New(notis.GenericsErrorWarnMsg)
	} else {
		role = rawRole
	}

	if id, ok := claims["userId"].(string); id == "" || !ok {
		return "", "", helper.GetPrimitiveTime(), errors.New(notis.GenericsErrorWarnMsg)
	} else {
		userId = id
	}

	if expPeriod, ok := claims["exp"].(time.Time); expPeriod == helper.GetPrimitiveTime() || !ok {
		return "", "", helper.GetPrimitiveTime(), errors.New(notis.GenericsErrorWarnMsg)
	} else {
		exp = expPeriod
	}

	return userId, role, exp, nil
}
