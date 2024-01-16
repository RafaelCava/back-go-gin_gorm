// domain/usecases/auth_usecase/auth_service.go - Camada de Domínio
package auth_usecase

import (
	encrypter_adapter "github.com/RafaelCava/kitkit-back-go/cmd/infra/encrypter_adapter"
	user_repository "github.com/RafaelCava/kitkit-back-go/cmd/infra/repositories/user_repository"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Validate(authRequest *AuthLoginRequest) (bool, error)
	GenerateToken(data interface{}) (string, error)
}

type AuthServiceImpl struct {
	userRepository   user_repository.UserRepository
	encrypterAdapter encrypter_adapter.EncrypterAdapter
}

func NewAuthServiceImpl(userRepository user_repository.UserRepository, encrypterAdapter encrypter_adapter.EncrypterAdapter) *AuthServiceImpl {
	return &AuthServiceImpl{
		userRepository,
		encrypterAdapter,
	}
}

func (s *AuthServiceImpl) Validate(authRequest *AuthLoginRequest) (bool, error) {
	user, err := s.userRepository.FindByEmail(authRequest.Email)
	if err != nil {
		return false, err
	}
	if user == nil {
		return false, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(authRequest.Password))
	if err == nil {
		return true, err
	} else {
		return false, nil
	}
}

func (s *AuthServiceImpl) GenerateToken(data interface{}) (string, error) {
	token, err := s.encrypterAdapter.GenerateToken(data)
	if err != nil {
		return "", err
	}
	return token, nil
}
