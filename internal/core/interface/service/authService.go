package service

type AuthService interface {
	Auth(login, password string) (bool, error)
	HashPassword(password string) (string, error)
}
