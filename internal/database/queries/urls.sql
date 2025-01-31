-- name: ListUrls :many
  SELECT * FROM urls;

-- name: GetUrlByShortUrl :one
  SELECT * FROM urls WHERE short_url = $1;

-- name: CreateUrl :one
  INSERT INTO urls (url, short_url, expired_at) VALUES ($1, $2, $3) RETURNING *;

-- name: DeleteUrlByShortUrl :one
  DELETE FROM urls WHERE short_url = $1 RETURNING *;
