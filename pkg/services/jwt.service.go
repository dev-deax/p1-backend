package service

import (
	"os"
	"p1-backend/api/pkg/models"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(user *models.Usuario) (models.Token, error) {
	expiration := time.Now().Add(time.Hour * 8)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":      user.Email,
		"name":       user.Nombre,
		"apellido":   user.Apellido,
		"rol_id":     user.RolID,
		"id":         strconv.Itoa(int(user.ID)),
		"expiration": expiration.Unix(),
	})
	secret := os.Getenv("JWT_SECRET")
	access_token, err := token.SignedString([]byte(secret))
	if err != nil {
		println("******" + err.Error() + "*******")
		return models.Token{}, err
	}
	return models.Token{Access_Token: access_token, Expiration: expiration}, nil
}

func ValidateToken(access_token string) (jwt.Claims, error) {
	signingKey := []byte(os.Getenv("JWT_SECRET"))
	token, err := jwt.Parse(access_token, func(token *jwt.Token) (interface{}, error) {
		return signingKey, nil
	})
	if err != nil {
		return nil, err
	}
	return token.Claims, err
}
