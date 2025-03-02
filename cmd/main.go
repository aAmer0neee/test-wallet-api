package main

import (
	"database/sql"
	"log"
	"time"

	"github.com/aAmer0neee/test-wallet-api/pkg/domain"
	"github.com/aAmer0neee/test-wallet-api/pkg/repository"
	"github.com/aAmer0neee/test-wallet-api/pkg/server"
	"github.com/aAmer0neee/test-wallet-api/pkg/service"
	_ "github.com/lib/pq"
)

var (
	port = ":8080"
)

func main() {

	time.Sleep(10 * time.Second)
	     
	database, err := sql.Open("postgres", "postgres://postgres:postgres@postgres:5432/postgres?sslmode=disable")
	
	database.SetMaxOpenConns(30)
	database.SetMaxIdleConns(5) 
	database.SetConnMaxLifetime(30 * time.Minute) 

	if err != nil {
		log.Fatalf("error opening database %v", err.Error())
	}
	defer database.Close()

	if err := database.Ping(); err != nil {
		log.Fatalf("error ping database %v", err.Error())

	}


	cacheWallets := domain.InitCache()

	repo := repository.InitRepository(database, cacheWallets)

	service := service.InitService(repo)

	server := server.InitServer(service)

	server.Up(port)

}
