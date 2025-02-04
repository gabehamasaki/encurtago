-- name: ListUrls :many
  SELECT * FROM urls order by created_at desc limit 10;

-- name: GetUrlByShortUrl :one
  SELECT * FROM urls WHERE short_url = $1;

-- name: CreateUrl :one
  INSERT INTO urls (url, short_url, expired_at) 
  VALUES ($1, $2, $3) 
  RETURNING *;

-- name: DeleteUrlByShortUrl :one
  DELETE FROM urls WHERE short_url = $1 RETURNING *;

-- name: UpdateClickCount :exec
  UPDATE urls SET click_count = click_count + 1 WHERE short_url = $1;
