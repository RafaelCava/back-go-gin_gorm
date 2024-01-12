// domain_layer/models/user_models/user.go - Camada de Domínio
package user_models

import "time"

// User representa um modelo de usuário.
type User struct {
	ID        string    `gorm:"primaryKey" json:"id"`
	Name      *string   `json:"name"`
	Email     string    `gorm:"unique;not null" json:"email"`
	Password  string    `gorm:"not null" json:"password"`
	CreatedAt time.Time `gorm:"index" json:"created_at"`
	UpdatedAt time.Time `gorm:"index" json:"updated_at"`
}
