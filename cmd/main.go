package main

import (
	"github.com/gin-gonic/gin"
	"github.com/novinbukannopin/fc-simple-forum/internal/handlers/memberships"
)

func main() {
	r := gin.Default()

	membershipHandler := memberships.Handler{r}
	membershipHandler.RegisterRoute()

	r.Run(":9999")
}
