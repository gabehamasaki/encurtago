package dtos

import "github.com/gabehamasaki/encurtago/internal/database"

type URL struct {
	ID        string `json:"id"`
	Original  string `json:"original"`
	Shortened string `json:"shortened"`
	CreatedAt string `json:"created_at"`
}

func (u *URL) ToDTO(raw *database.Url) {
	u.ID = raw.ID.String()
	u.Original = raw.Url
	u.Shortened = raw.ShortUrl
	u.CreatedAt = raw.CreatedAt.Time.String()
}

type CreateShortURLRequest struct {
	Original string `json:"original"`
}

type CreateShortURLResponse struct {
	Original  string `json:"original"`
	Shortened string `json:"shortened"`
}

func (r *CreateShortURLResponse) ToDTO(raw *database.Url) {
	r.Original = raw.Url
	r.Shortened = raw.ShortUrl
}
