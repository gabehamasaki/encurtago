package main

import (
	"github.com/gabehamasaki/encurtago/client"
	"github.com/gabehamasaki/encurtago/internal/router"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	client.RegisterHandlers(r)
	router.RegisterRoutes(r)

	r.Run(":8080")
}
