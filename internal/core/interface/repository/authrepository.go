package repository

type AuthRepository interface {
	Auth(login, password string) (bool, error)
}
