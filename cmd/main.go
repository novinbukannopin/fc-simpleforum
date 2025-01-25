package main

import (
	"github.com/gin-gonic/gin"
	"github.com/novinbukannopin/fc-simple-forum/internal/configs"
	"github.com/novinbukannopin/fc-simple-forum/internal/handlers/memberships"
	"log"
)

func main() {
	r := gin.Default()

	var (
		cfg *configs.Config
	)

	err := configs.Init(
		configs.WithConfigFolder([]string{"./internal/configs"}),
		configs.WithConfigFile("config"),
		configs.WithConfigType("yaml"),
	)

	if err != nil {
		log.Fatal("failed to load config")
	}

	cfg = configs.Get()
	log.Println("config loaded", cfg)

	membershipHandler := memberships.Handler{Engine: r}
	membershipHandler.RegisterRoute()

	r.Run(cfg.Service.Port)
}
