package storage

import (
	"errors"

	"github.com/WeisseNacht18/gophermart/internal/builder"
)

type JWTStorage struct {
	Tokens map[string]string
}

var jwtStorage JWTStorage

func NewJWTStorage() {
	jwtStorage = JWTStorage{
		Tokens: make(map[string]string),
	}
}

func AddToken(login string) (token string, err error) {
	token, err = builder.BuildJWTStringWithLogin(login)

	if err == nil {
		jwtStorage.Tokens[token] = login
	}
	return
}

func FindToken(token string) (login string, err error) {
	login, ok := jwtStorage.Tokens[token]
	if !ok {
		err = errors.New("token was not found")
	}
	return
}

func DeleteToken() {

}
