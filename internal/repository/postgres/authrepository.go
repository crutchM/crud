package postgres

import "crud/internal/lib/db"

type AuthRepository struct {
	*db.DB
}

func (auth AuthRepository) Auth(login, password string) (bool, error) {
	return false, nil
}
