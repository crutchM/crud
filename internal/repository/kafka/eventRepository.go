package kafka

import (
	"context"
	"crud/internal/core/interface/repository"
	"crud/internal/core/model"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/segmentio/kafka-go"
	"log"
)

type _eventRepository struct {
	conn *kafka.Conn
}

func NewEventRepo(host string) repository.EventRepository {
	conn, err := kafka.DialLeader(context.Background(), "tcp", host, "post", 0)

	if err != nil {
		log.Fatal(err)
	}

	return _eventRepository{conn: conn}

}

func (repo _eventRepository) SendEvent(ctx context.Context, event model.Event) error {
	bytes, err := json.Marshal(event)

	if err != nil {
		return errors.New("error marshaling message")
	}

	_, err = repo.conn.WriteMessages(
		kafka.Message{
			Topic: "post",
			Key:   []byte(event.Id),
			Value: bytes,
		})

	if err != nil {
		return fmt.Errorf("error sending message %s", err.Error())
	}

	return err
}
