package handlers

import (
	"github.com/gabehamasaki/encurtago/internal/dtos"
	"github.com/gin-gonic/gin"
)

func (h *Handler) ListAllURLs(ctx *gin.Context) {
	urlsRaw, err := h.cfg.DB.ListUrls(ctx)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	urls := []dtos.URL{}
	for _, raw := range urlsRaw {
		var url dtos.URL
		url.ToDTO(&raw)
		urls = append(urls, url)
	}

	ctx.JSON(200, gin.H{"urls": urls})
}
