package service

import (
	"context"
	"crud/internal/core/interface/repository"
	"crud/internal/core/interface/service"
	"crud/internal/core/model"
	"errors"
	"fmt"
	"github.com/hashicorp/go-uuid"
	"log/slog"
)

type _postService struct {
	repo  repository.PostRepository
	kafka repository.EventRepository
}

func NewPostService(repo repository.PostRepository, kafka repository.EventRepository) service.PostService {
	return _postService{repo: repo, kafka: kafka}
}

func (postService _postService) CreatePost(ctx context.Context, post model.Post) (int, error) {
	id, err := postService.repo.CreatePost(ctx, post)

	if err != nil {
		slog.Error(err.Error())
		return 0, errors.New("ошибка создания поста")
	}

	requestId, err := uuid.GenerateUUID()

	if err != nil {
		slog.Error("ошибка генерации uuid: ", err.Error())
		return id, nil
	}
	event := model.Event{
		Id:    requestId,
		Key:   requestId,
		Value: id,
	}

	err = postService.kafka.SendEvent(ctx, event)

	if err != nil {
		slog.Error(fmt.Sprint("ошибка отправки сообщения: ", err.Error()))
		return id, nil
	}

	return id, nil
}

func (postService _postService) GetPost(ctx context.Context, postId int) (model.Post, error) {
	return postService.repo.GetPost(ctx, postId)
}
