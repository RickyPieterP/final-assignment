package usecase

import (
	"mygram/app/usecase/request"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func GeneratePassword(pass []byte) (string, error) {
	hash, err := bcrypt.GenerateFromPassword(pass, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func ValidatePassword(hash, pass []byte) error {
	return bcrypt.CompareHashAndPassword(hash, pass)
}

func GenerateToken(in request.JWTToken) (string, error) {
	var secret = []byte("SecretYouShouldHide")
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(10 * time.Minute)
	claims["id"] = in.Id
	claims["user"] = in.Username
	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
