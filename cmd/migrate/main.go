package main

import (
	"fmt"

	"github.com/joqd/authify/internal/adapter/config"
	"github.com/joqd/authify/internal/adapter/storage/postgres"
)


func main() {
	conf := config.LoadConfig()
	db := postgres.NewPostgresDatabase(conf)
	
	if err := db.AutoMigrate(); err != nil {
		panic(err)
	}

	fmt.Println("migrate successfully done.")
}