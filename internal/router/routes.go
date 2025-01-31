package router

import (
	"github.com/gabehamasaki/encurtago/internal/config"
	"github.com/gabehamasaki/encurtago/internal/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, cfg *config.Config) {
	handler := handlers.NewHandler(cfg)

	api := r.Group("/api")
	{
		api.GET("/ping", handler.Ping)
		api.GET("/urls", handler.ListAllURLs)
		api.POST("/urls", handler.CreateShortURL)
	}

	r.GET("/r/:shortened", handler.Redirect)
}
