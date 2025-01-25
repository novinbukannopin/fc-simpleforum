package main

import (
	"github.com/gin-gonic/gin"
	"github.com/novinbukannopin/fc-simple-forum/internal/configs"
	"github.com/novinbukannopin/fc-simple-forum/internal/handlers/memberships"
	membershipsRepo "github.com/novinbukannopin/fc-simple-forum/internal/repository/memberships"
	"github.com/novinbukannopin/fc-simple-forum/pkg/internalsql"
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

	db, err := internalsql.Connect(cfg.Database.DataSourceName)
	if err != nil {
		log.Fatal("failed to connect to database")
	}

	_ = membershipsRepo.NewRepository(db)

	membershipHandler := memberships.Handler{Engine: r}
	membershipHandler.RegisterRoute()

	r.Run(cfg.Service.Port)
}
