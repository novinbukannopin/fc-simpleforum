package main

import (
	"github.com/gin-gonic/gin"
	"github.com/novinbukannopin/fc-simple-forum/internal/configs"
	membershipHandler "github.com/novinbukannopin/fc-simple-forum/internal/handlers/memberships"
	membershipRepo "github.com/novinbukannopin/fc-simple-forum/internal/repository/memberships"
	membershipService "github.com/novinbukannopin/fc-simple-forum/internal/service/memberships"
	"github.com/novinbukannopin/fc-simple-forum/pkg/internalsql"
	"log"
)

func main() {
	r := gin.Default()

	var (
		cfg *configs.Config
	)

	err := configs.Init(
		configs.WithConfigFolder([]string{"./internal/configs/"}),
		configs.WithConfigFile("config"),
		configs.WithConfigType("yaml"),
	)

	if err != nil {
		log.Fatal("failed to load config")
	}

	cfg = configs.Get()
	log.Println("config loaded", cfg)

	db, err := internalsql.Connect(cfg.Database.DataSourceName)
	if err != nil {
		log.Fatal("failed to connect to database")
	}

	membershipRepository := membershipRepo.NewRepository(db)
	membershipSvc := membershipService.NewService(cfg, membershipRepository)
	membershipHandlers := membershipHandler.NewHandler(r, membershipSvc)

	membershipHandlers.RegisterRoute()

	err = r.Run(cfg.Service.Port)
	if err != nil {
		return
	}
}
