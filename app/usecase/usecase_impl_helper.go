package usecase

import (
	"encoding/json"
	"fmt"
	"mygram/app/usecase/request"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type Token struct {
	Id       int
	Email    string
	Username string
}

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

	tok := Token{
		Email:    in.Email,
		Username: in.Username,
		Id:       in.Id,
	}
	claims := jwt.MapClaims{
		"exp":     time.Now().Add(3 * time.Minute).Unix(),
		"payload": tok,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ValidateToken(tokString string) (*Token, error) {
	var secret = "SecretYouShouldHide"
	errResp := fmt.Errorf("need signin")

	tok, err := jwt.Parse(tokString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errResp
		}
		return []byte(secret), nil
	})

	if err != nil {
		return nil, errResp
	}

	claims, ok := tok.Claims.(jwt.MapClaims)
	if !ok && !tok.Valid {
		return nil, errResp
	}

	payload, ok := claims["payload"]
	if !ok && !tok.Valid {
		return nil, errResp
	}

	bytePayload, err := json.Marshal(payload)
	if err != nil {
		fmt.Println(err)
	}

	var token Token
	err = json.Unmarshal(bytePayload, &token)
	if err != nil {
		fmt.Println(err)
	}

	return &token, nil
}
