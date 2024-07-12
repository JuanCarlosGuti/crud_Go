package main

import (
	"context"
	"fmt"
	"github.com/JuanCarlosGuti/Go_users.git/internal/user"
	"github.com/JuanCarlosGuti/Go_users.git/pkg/bootstrap"
	"github.com/JuanCarlosGuti/Go_users.git/pkg/handler"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main() {
	_ = godotenv.Load()
	server := http.NewServeMux()
	db, err := bootstrap.NewDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	logger := bootstrap.NewLogger()
	repo := user.NewRepository(db, logger)
	service := user.NewService(logger, repo)
	ctx := context.Background()
	handler.NewUserHttpServer(ctx, server, user.MakeEndpoints(ctx, service))
	port := os.Getenv("PORT")
	fmt.Println("Starting server at localhost: ", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), server))

}
