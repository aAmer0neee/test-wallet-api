package main

import (
	"database/sql"
	"log"

	"github.com/aAmer0neee/test-wallet-api/pkg/repository"
	"github.com/aAmer0neee/test-wallet-api/pkg/server"
	"github.com/aAmer0neee/test-wallet-api/pkg/service"
	_ "github.com/lib/pq"
)

var (
	port = ":8080"
)

func main() {
	database, err := sql.Open("postgres", "postgres://postgres:postgres@postgres:5432/postgres?sslmode=disable")

	if err != nil {
		log.Fatalf("error %v", err.Error())
	}
	defer database.Close()

	if err := database.Ping(); err != nil {
		log.Fatalf("error %v", err.Error())
	}

	repo := repository.InitRepository(database)

	service := service.InitService(repo)

	server := server.InitServer(service)

	server.Up(port)

}
