package repository

import (
	"crud/internal/repository/postgres"
)

type RepositoryManager struct {
	postgres.AuthRepository
}
