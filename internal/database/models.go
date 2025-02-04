// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0

package database

import (
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type Url struct {
	ID         uuid.UUID
	Url        string
	ShortUrl   string
	ClickCount pgtype.Int4
	CreatedAt  pgtype.Timestamptz
	ExpiredAt  time.Time
}
