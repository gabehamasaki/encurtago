package handlers

import (
	"crypto/sha256"
	"fmt"
	"math/big"
	"math/rand"
	"time"

	"github.com/gabehamasaki/encurtago/internal/database"
	"github.com/gabehamasaki/encurtago/internal/dtos"
	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateShortURL(ctx *gin.Context) {
	var req dtos.CreateShortURLRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	expiredAt := time.Now().Add(7 * 24 * time.Hour) // Adds 7 days

	url, err := h.cfg.DB.CreateUrl(ctx, database.CreateUrlParams{
		Url:       req.Original,
		ShortUrl:  makeShortUrl(req.Original),
		ExpiredAt: expiredAt,
	})
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	var res dtos.CreateShortURLResponse
	res.ToDTO(&url)
	ctx.JSON(200, gin.H{"url": res})
}

const base62Chars = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

func makeShortUrl(url string) string {
	rd := rand.New(rand.NewSource(time.Now().UnixNano()))

	salt := rd.Intn(10000)

	input := fmt.Sprintf("%s-%d", url, salt)

	hash := sha256.Sum256([]byte(input))

	bigInt := new(big.Int).SetBytes(hash[:])

	shortUrl := encodeBase62(bigInt)

	return shortUrl[:8]
}

func encodeBase62(num *big.Int) string {
	base := big.NewInt(62)
	zero := big.NewInt(0)
	result := ""

	for num.Cmp(zero) > 0 {
		mod := new(big.Int)
		num.DivMod(num, base, mod)
		result = string(base62Chars[mod.Int64()]) + result
	}

	return result
}
