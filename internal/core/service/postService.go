package service

import (
	"context"
	"crud/internal/core/interface/repository"
	"crud/internal/core/interface/service"
	"crud/internal/core/model"
	"errors"
	"log/slog"
)

type _postService struct {
	repo repository.PostRepository
}

func NewPostService(repo repository.PostRepository) service.PostService {
	return _postService{repo: repo}
}

func (postService _postService) CreatePost(ctx context.Context, post model.Post) (int, error) {
	id, err := postService.repo.CreatePost(ctx, post)

	if err != nil {
		slog.Error(err.Error())
		return 0, errors.New("ошибка создания поста")
	}

	return id, nil
}

func (postService _postService) GetPost(ctx context.Context, postId int) (model.Post, error) {
	return postService.repo.GetPost(ctx, postId)
}
