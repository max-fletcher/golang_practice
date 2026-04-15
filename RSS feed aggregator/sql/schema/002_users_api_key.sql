-- +goose Up
ALTER TABLE users ADD COLUMN api_key VARCHAR(64) UNIQUE NOT NULL DEFAULT (
    encode(sha256(random()::text::bytea), 'hex') -- creates a random unique value to store in api_key field
);

-- +goose Down
ALTER TABLE users DROP COLUMN api_key;