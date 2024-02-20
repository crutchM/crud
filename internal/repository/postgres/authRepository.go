package postgres

import (
	"crud/internal/core/interface/repository"
	"crud/internal/lib/db"
	"fmt"
)

type userDB struct {
	Login    string `db:"login"`
	Password string `db:"password"`
}

type _authRepo struct {
	*db.Db
}

func NewRepo(db *db.Db) repository.AuthRepository {
	return _authRepo{db}
}

func (repo _authRepo) GetUser(login, hashPassword string) (string, error) {
	var user userDB

	row := repo.PgConn.QueryRow(nil, `SELECT * FROM user WHERE login=$1, password=$2`, login, hashPassword)

	if err := row.Scan(&user); err != nil {
		return "", fmt.Errorf("не смогли получить юзера: %x", err)
	}

	return login, nil

}

func (repo _authRepo) Register(login, hashPassword string) (string, error) {

	_, err := repo.PgConn.Exec(
		nil,
		`INSERT INTO user(login, password) values ($1, $2)`,
		login, hashPassword,
	)

	if err != nil {
		return "", fmt.Errorf("не смогли создать: %x", err)
	}

	return login, nil
}
