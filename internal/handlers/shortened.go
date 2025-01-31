package handlers

import "github.com/gin-gonic/gin"

func (h *Handler) Redirect(ctx *gin.Context) {
	shortened := ctx.Param("shortened")
	url, err := h.cfg.DB.GetUrlByShortUrl(ctx, shortened)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.Redirect(301, url.Url)
}
