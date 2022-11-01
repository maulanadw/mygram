package helpers

import (
	"errors"
	"mygram/configs"

	"github.com/golang-jwt/jwt/v4"
)

func GenerateToken(ID int, email string) (string, error) {
	claims := jwt.MapClaims{
		"id":    ID,
		"email": email,
	}

	parseToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := parseToken.SignedString([]byte(configs.App.SecretKey))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func VerifyToken(tokenStr string) (interface{}, error) {
	errResp := errors.New("token invalid")
	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errResp
		}

		return []byte(configs.App.SecretKey), nil
	})

	if err != nil {
		return nil, err
	}

	if _, ok := token.Claims.(jwt.MapClaims); !ok && !token.Valid {
		return nil, errResp
	}

	return token.Claims.(jwt.MapClaims), nil
}
