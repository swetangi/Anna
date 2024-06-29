package jwt

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateJwt(email string) (string, error) {
	fmt.Println("email in gen :", email)
	expirationTime := time.Now().Add(2 * time.Minute)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.MapClaims{
		"expire": expirationTime.Unix(),
		"email":  email,
	})

	secretKey := "abc"
	secretKeyByte := []byte(secretKey)
	tokenStr, err := token.SignedString(secretKeyByte)
	if err != nil {
		return "", err
	}
	fmt.Println("tokenStr ", tokenStr)
	return tokenStr, nil
}

func VerifyToken(tokenStr string) (jwt.Token, error) {
	fmt.Println("token STr : ", tokenStr)
	secretKey := "abc"
	secretKeyByte := []byte(secretKey)

	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		return secretKeyByte, nil
	})
	if err != nil {
		return jwt.Token{}, err
	}

	if !token.Valid {
		return jwt.Token{}, fmt.Errorf("invalid token")
	}
	return *token, nil
}
