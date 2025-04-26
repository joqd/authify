package main

import (
	"github.com/joqd/authify/internal/adapter/config"
	"github.com/joqd/authify/internal/adapter/server"
	"github.com/joqd/authify/internal/adapter/storage/postgres"
)


func main() {
	conf := config.LoadConfig()
	db := postgres.NewPostgresDatabase(conf)
	server := server.NewServer(conf, db)

	server.Start()
}