package postgres

import (
	"context"
	"crud/internal/core/interface/repository"
	"crud/internal/core/model"
	"crud/internal/lib/db"
	"crud/internal/repository/dbModel"
	"fmt"
)

type _postRepository struct {
	db *db.Db
}

func NewPostRepo(db *db.Db) repository.PostRepository {
	return _postRepository{db}
}

func (postRepository _postRepository) CreatePost(ctx context.Context, post model.Post) (int, error) {
	postDb := dbModel.Post(post)
	var id int

	err := postRepository.db.PgConn.QueryRow(ctx,
		`INSERT INTO public.post(title, body, image, author) values ($1,$2,$3,$4) RETURNING id`,
		postDb.Title,
		postDb.Body,
		postDb.ImageURL,
		postDb.Author).Scan(&id)

	return id, err
}

func (postRepository _postRepository) GetPost(ctx context.Context, postId int) (model.Post, error) {
	var post dbModel.Post

	err := postRepository.db.PgConn.QueryRow(ctx,
		`SELECT p.title, p.body, p.image, p.author FROM public.post p WHERE p.id=$1`,
		postId).Scan(&post.Title, &post.Body, &post.ImageURL, &post.Author)

	if err != nil {
		return model.Post{}, fmt.Errorf("ошибка получения поста: %s", err.Error())
	}

	return model.Post(post), nil

}
