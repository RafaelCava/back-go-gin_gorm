package encrypter_adapter

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var secret = []byte(os.Getenv("SECRET"))

type EncrypterAdapter interface {
	GenerateToken(data interface{}) (string, error)
}

type EncrypterAdapterImpl struct{}

func NewEncrypterAdapterImpl() *EncrypterAdapterImpl {
	return &EncrypterAdapterImpl{}
}

func (e *EncrypterAdapterImpl) GenerateToken(data interface{}) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["data"] = data
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
