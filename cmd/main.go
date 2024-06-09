package main

import (
	"context"
	"fmt"
	"github.com/JuanCarlosGuti/Go_users.git/internal/user"
	"github.com/JuanCarlosGuti/Go_users.git/pkg/bootstrap"
	"log"
	"net/http"
)

func main() {
	server := http.NewServeMux()
	db := bootstrap.NewDB()
	logger := bootstrap.NewLogger()
	repo := user.NewRepository(db, logger)
	service := user.NewService(logger, repo)
	ctx := context.Background()

	server.HandleFunc("/users", user.MakeEndpoints(ctx, service))
	fmt.Println("Starting server at localhost:8090")
	log.Fatal(http.ListenAndServe(":8090", server))

}
