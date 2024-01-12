// infra/user_repository/user_repository.go - Camada de Dados (Data Layer)
package user_repository

import (
	"github.com/RafaelCava/kitkit-back-go/cmd/domain/models/user_models"
	"gorm.io/gorm"
)

// UserRepository é uma interface que define operações de banco de dados relacionadas ao usuário.
type UserRepository interface {
	FindByID(userID string) (*user_models.User, error)
	Find() ([]*user_models.User, error)
	Create(user *user_models.User) (string, error)
}

// GormUserRepository implementa UserRepository usando Gorm.
type GormUserRepository struct {
	db *gorm.DB
}

// NewGormUserRepository cria uma nova instância de GormUserRepository.
func NewGormUserRepository(db *gorm.DB) *GormUserRepository {
	return &GormUserRepository{db}
}

// FindByID recupera um usuário por ID do banco de dados usando Gorm.
func (r *GormUserRepository) FindByID(userID string) (*user_models.User, error) {
	var user user_models.User
	if err := r.db.First(&user, "id = ?", userID).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *GormUserRepository) Create(user *user_models.User) (string, error) {
	if err := r.db.Create(&user).Error; err != nil {
		return "", err
	}
	return user.ID, nil
}

func (r *GormUserRepository) Find() ([]*user_models.User, error) {
	var users []*user_models.User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
