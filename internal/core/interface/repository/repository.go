package repository

type AuthRepository interface {
	GetUser(login, hashPassword string) (string, error)
	Register(login, hashPassword string) (string, error)
}
