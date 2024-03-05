package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/segmentio/kafka-go"
	"log"
	"log/slog"
)

type Message struct {
	Id    string `json:"Id"`
	Key   string `json:"Key"`
	Value int    `json:"Value"`
}

type Post struct {
	Title    string `db:"title"`
	Body     string `db:"body"`
	ImageURL string `db:"image"`
	Author   string `db:"author"`
}

func main() {

	ctx := context.Background()

	db := New(ctx)
	r := kafka.ReaderConfig{
		Brokers:   []string{"10.80.0.139:29092"},
		Topic:     "post",
		Partition: 0,
		MaxBytes:  10e6,
	}
	conn := kafka.NewReader(r)

	//err = conn.SetReadDeadline(time.Now().Add(10 * time.Second))
	//
	//if err != nil {
	//	log.Fatal(err)
	//}
	defer func() {
		if err := conn.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	//defer func() {
	//	if err := batch.Close(); err != nil {
	//		log.Fatal(err)
	//	}
	//}()
	for {
		msg, err := conn.ReadMessage(context.Background())

		if err != nil {
			slog.Error(fmt.Sprint("error receiving message: ", err.Error()))
			continue
		}
		var message Message

		err = json.Unmarshal(msg.Value, &message)

		if err != nil {
			slog.Error("error decoding message")
			continue
		}
		post, err := getPost(ctx, db, message.Value)

		if err != nil {
			slog.Error(fmt.Sprint("error creating post: ", err.Error()))
		}
		fmt.Println(msg.Offset)
		fmt.Println(string(msg.Key))
		fmt.Println(post)

		SendMessage(post)

		conn.SetOffset(msg.Offset + 1)
	}
}

func getPost(ctx context.Context, db *Db, id int) (Post, error) {
	var post Post

	err := db.PgConn.QueryRow(ctx,
		`SELECT p.title, p.body, p.image, p.author FROM public.post p WHERE p.id=$1`,
		id).Scan(&post.Title, &post.Body, &post.ImageURL, &post.Author)

	if err != nil {
		return Post{}, fmt.Errorf("ошибка получения поста: %s", err.Error())
	}

	return post, nil
}
