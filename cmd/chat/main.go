package main

import (
	api "github.com/davidPardoC/go-chat/cmd/chat/api/http"
	"github.com/davidPardoC/go-chat/internal/config"
)

func main() {
	cfg := config.LoadConfig()
	config.ConectDatabase(cfg)
	api.StartHttpServer()
}
