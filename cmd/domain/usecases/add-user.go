package domain

type AddUser interface {
	AddUser(username, password string) (int, error)
}
