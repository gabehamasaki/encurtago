package main

import (
	"github.com/gabehamasaki/encurtago/client"
	"github.com/gabehamasaki/encurtago/internal/config"
	"github.com/gabehamasaki/encurtago/internal/router"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	cfg := config.NewConfig()

	client.RegisterHandlers(r, cfg.ENV)
	router.RegisterRoutes(r)

	r.Run(":8080")
}
