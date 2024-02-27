package postgres

import (
	"context"
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

func (repo _authRepo) GetUser(ctx context.Context, login, hashPassword string) (string, error) {
	var user userDB

	row := repo.PgConn.QueryRow(ctx, `SELECT * FROM public.user WHERE login=$1 AND pas=$2`, login, hashPassword)

	if err := row.Scan(&user); err != nil {
		return "", fmt.Errorf("не смогли получить юзера: %x", err)
	}

	return login, nil

}

func (repo _authRepo) Register(ctx context.Context, login, hashPassword string) (string, error) {
	_, err := repo.PgConn.Exec(
		ctx,
		`INSERT INTO public.user(login, pass) values ($1, $2)`,
		login, hashPassword,
	)

	if err != nil {
		return "", fmt.Errorf("не смогли создать: %x", err)
	}

	return login, nil
}
