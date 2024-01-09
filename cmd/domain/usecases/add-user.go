package domain

import "github.com/RafaelCava/kitkit-back-go/cmd/domain/models"

type AddUser interface {
	Create(user models.User) (id string, err error)
}

type UserStore struct{}

func (us *UserStore) Create(user models.User) (id string, err error) {
	return
}
