package main

import (
	"log"

	api "github.com/davidPardoC/go-chat/cmd/chat/api/http"
	"github.com/davidPardoC/go-chat/internal/config"
)

func main() {
	cfg := config.LoadConfig()

	database, err := config.ConectDatabase(cfg)

	config.AutomigrateDatabase(database)

	if err != nil {
		log.Fatal(err.Error())
	}

	app := api.NewServerApp(database)

	app.StartHttpServer()
}
