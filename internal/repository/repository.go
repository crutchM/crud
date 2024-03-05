package repository

import (
	"crud/internal/core/interface/repository"
	"crud/internal/lib/db"
	"crud/internal/repository/kafka"
	"crud/internal/repository/postgres"
)

type RepositoryManager struct {
	repository.AuthRepository
	repository.PostRepository
	repository.EventRepository
}

func NewRepositoryManager(db *db.Db, host string) RepositoryManager {
	return RepositoryManager{
		postgres.NewRepo(db),
		postgres.NewPostRepo(db),
		kafka.NewEventRepo(host),
	}
}
