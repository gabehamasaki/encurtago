package main

import (
	"context"
	"fmt"

	"github.com/gabehamasaki/encurtago/client"
	"github.com/gabehamasaki/encurtago/internal/config"
	"github.com/gabehamasaki/encurtago/internal/database"
	"github.com/gabehamasaki/encurtago/internal/database/connection"
	"github.com/gabehamasaki/encurtago/internal/router"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	cfg := config.NewConfig()

	conn, err := connection.NewConnection(context.Background(), cfg)
	if err != nil {
		panic(err)
	}
	defer conn.Close(context.Background())

	db := database.New(conn)
	cfg.SetDB(db)

	client.RegisterHandlers(r, cfg.ENV)
	router.RegisterRoutes(r, cfg)

	r.Run(fmt.Sprintf(":%s", cfg.PORT))
}
