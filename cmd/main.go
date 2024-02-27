package main

import (
	"context"
	"crud/internal/core/service"
	"crud/internal/lib/db"
	"crud/internal/repository"
	"crud/internal/transport/http"
	"log"
	http2 "net/http"
	"time"
)

func main() {

	timeout := time.Second * 10

	ctx := context.Background()

	withTimeout, _ := context.WithTimeout(ctx, timeout)

	database := db.New(withTimeout)

	manager := repository.NewRepositoryManager(database)

	serv := service.NewAuthService(manager.AuthRepository)

	postServ := service.NewPostService(manager.PostRepository)

	router := http.InitRoutes(serv, postServ)

	if err := http2.ListenAndServe(":2222", router); err != nil {
		log.Fatal(err)
	}
}
