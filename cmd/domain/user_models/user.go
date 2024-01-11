// domain_layer/user_models/user.go - Camada de Domínio
package user_models

// User representa um modelo de usuário.
type User struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Name string `json:"name"`
	// outros campos do usuário...
}
