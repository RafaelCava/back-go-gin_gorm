// infra/user_repository/user_repository.go - Camada de Dados (Data Layer)
package user_repository

import (
	"github.com/RafaelCava/kitkit-back-go/domain/models/user_models"
	"gorm.io/gorm"
)

// UserRepository é uma interface que define operações de banco de dados relacionadas ao usuário.
type UserRepository interface {
	FindByID(userID string) (*user_models.User, error)
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
	if err := r.db.First(&user, userID).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
