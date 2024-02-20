package service

type AuthService interface {
	Register(login, password string) (string, error)
	GenerateToken(login, password string) (string, error)
}
