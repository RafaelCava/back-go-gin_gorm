// infra/repository/user_repository/user_repository.go - Camada de Dados (Data Layer)
package user_repository

import (
	"github.com/RafaelCava/kitkit-back-go/cmd/domain/models/user_models"
	"gorm.io/gorm"
)

// UserRepository é uma interface que define operações de banco de dados relacionadas ao usuário.
type UserRepository interface {
	FindByID(userID string) (*user_models.UserWithoutPassword, error)
	FindByEmail(email string) (*user_models.User, error)
	Find() ([]*user_models.UserWithoutPassword, error)
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
func (r *GormUserRepository) FindByID(userID string) (*user_models.UserWithoutPassword, error) {
	var user user_models.User
	if err := r.db.First(&user, "id = ?", userID).Error; err != nil {
		return nil, err
	}

	return &user_models.UserWithoutPassword{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}, nil
}

func (r *GormUserRepository) Create(user *user_models.User) (string, error) {
	if err := r.db.Create(&user).Error; err != nil {
		return "", err
	}
	return user.ID, nil
}

func (r *GormUserRepository) Find() ([]*user_models.UserWithoutPassword, error) {
	var users []*user_models.User
	var usersWithoutPassword []*user_models.UserWithoutPassword
	if err := r.db.Select([]string{"id", "name", "email", "created_at", "updated_at"}).Find(&users).Error; err != nil {
		return nil, err
	}
	for _, user := range users {
		usersWithoutPassword = append(usersWithoutPassword, &user_models.UserWithoutPassword{
			ID:        user.ID,
			Name:      user.Name,
			Email:     user.Email,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		})
	}
	return usersWithoutPassword, nil
}

func (r *GormUserRepository) FindByEmail(email string) (*user_models.User, error) {
	var user user_models.User
	if err := r.db.First(&user, "email = ?", email).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
