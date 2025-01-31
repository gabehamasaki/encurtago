package router

import (
	"github.com/gabehamasaki/encurtago/internal/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.GET("/ping", handlers.Ping)
	}
}
