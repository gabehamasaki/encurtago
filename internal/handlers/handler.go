package handlers

import "github.com/gabehamasaki/encurtago/internal/config"

type Handler struct {
	cfg *config.Config
}

func NewHandler(cfg *config.Config) *Handler {
	return &Handler{cfg: cfg}
}
