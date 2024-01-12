// domain/usecases/user_usecase/user_service.go - Camada de Domínio
package user_usecase

import (
	"github.com/RafaelCava/kitkit-back-go/cmd/domain/models/user_models"
	infra "github.com/RafaelCava/kitkit-back-go/cmd/infra/repositories/user_repository"
)

// UserService é uma interface que define operações relacionadas ao usuário.
type UserService interface {
	GetUserByID(userID string) (*user_models.User, error)
	CreateUser(user *user_models.User) (string, error)
	FindAll() ([]*user_models.User, error)
}

// UserServiceImpl implementa UserService usando um UserRepository.
type UserServiceImpl struct {
	userRepository infra.UserRepository
}

// NewUserServiceImpl cria uma nova instância de UserServiceImpl.
func NewUserServiceImpl(userRepository infra.UserRepository) *UserServiceImpl {
	return &UserServiceImpl{userRepository}
}

// GetUserByID recupera um usuário por ID usando o UserRepository.
func (s *UserServiceImpl) GetUserByID(userID string) (*user_models.User, error) {
	return s.userRepository.FindByID(userID)
}

func (s *UserServiceImpl) CreateUser(user *user_models.User) (string, error) {
	return s.userRepository.Create(user)
}

func (s *UserServiceImpl) FindAll() ([]*user_models.User, error) {
	return s.userRepository.Find()
}
