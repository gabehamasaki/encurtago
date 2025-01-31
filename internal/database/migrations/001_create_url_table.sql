-- Write your migrate up statements here
CREATE TABLE IF NOT EXISTS urls (
    "id" uuid PRIMARY KEY NOT NULL DEFAULT  gen_random_uuid(),
    "url" VARCHAR(255) NOT NULL,
    "short_url" VARCHAR(32) NOT NULL,
    "click_count" INT DEFAULT 0,
    "created_at" TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    "expired_at" TIMESTAMPTZ NOT NULL
);

---- create above / drop below ----
DROP TABLE IF EXISTS urls;

-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.
