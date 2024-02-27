package service

import (
	"context"
	"crud/internal/core/model"
)

type AuthService interface {
	Register(ctx context.Context, login, password string) (string, error)
	GenerateToken(ctx context.Context, login, password string) (string, error)
}

type PostService interface {
	CreatePost(ctx context.Context, post model.Post) (int, error)
	GetPost(ctx context.Context, postId int) (model.Post, error)
}
