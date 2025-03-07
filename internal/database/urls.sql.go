// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: urls.sql

package database

import (
	"context"
	"time"
)

const createUrl = `-- name: CreateUrl :one
  INSERT INTO urls (url, short_url, expired_at) 
  VALUES ($1, $2, $3) 
  RETURNING id, url, short_url, click_count, created_at, expired_at
`

type CreateUrlParams struct {
	Url       string
	ShortUrl  string
	ExpiredAt time.Time
}

func (q *Queries) CreateUrl(ctx context.Context, arg CreateUrlParams) (Url, error) {
	row := q.db.QueryRow(ctx, createUrl, arg.Url, arg.ShortUrl, arg.ExpiredAt)
	var i Url
	err := row.Scan(
		&i.ID,
		&i.Url,
		&i.ShortUrl,
		&i.ClickCount,
		&i.CreatedAt,
		&i.ExpiredAt,
	)
	return i, err
}

const deleteUrlByShortUrl = `-- name: DeleteUrlByShortUrl :one
  DELETE FROM urls WHERE short_url = $1 RETURNING id, url, short_url, click_count, created_at, expired_at
`

func (q *Queries) DeleteUrlByShortUrl(ctx context.Context, shortUrl string) (Url, error) {
	row := q.db.QueryRow(ctx, deleteUrlByShortUrl, shortUrl)
	var i Url
	err := row.Scan(
		&i.ID,
		&i.Url,
		&i.ShortUrl,
		&i.ClickCount,
		&i.CreatedAt,
		&i.ExpiredAt,
	)
	return i, err
}

const getUrlByShortUrl = `-- name: GetUrlByShortUrl :one
  SELECT id, url, short_url, click_count, created_at, expired_at FROM urls WHERE short_url = $1
`

func (q *Queries) GetUrlByShortUrl(ctx context.Context, shortUrl string) (Url, error) {
	row := q.db.QueryRow(ctx, getUrlByShortUrl, shortUrl)
	var i Url
	err := row.Scan(
		&i.ID,
		&i.Url,
		&i.ShortUrl,
		&i.ClickCount,
		&i.CreatedAt,
		&i.ExpiredAt,
	)
	return i, err
}

const listUrls = `-- name: ListUrls :many
  SELECT id, url, short_url, click_count, created_at, expired_at FROM urls order by created_at desc limit 10
`

func (q *Queries) ListUrls(ctx context.Context) ([]Url, error) {
	rows, err := q.db.Query(ctx, listUrls)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Url
	for rows.Next() {
		var i Url
		if err := rows.Scan(
			&i.ID,
			&i.Url,
			&i.ShortUrl,
			&i.ClickCount,
			&i.CreatedAt,
			&i.ExpiredAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateClickCount = `-- name: UpdateClickCount :exec
  UPDATE urls SET click_count = click_count + 1 WHERE short_url = $1
`

func (q *Queries) UpdateClickCount(ctx context.Context, shortUrl string) error {
	_, err := q.db.Exec(ctx, updateClickCount, shortUrl)
	return err
}
