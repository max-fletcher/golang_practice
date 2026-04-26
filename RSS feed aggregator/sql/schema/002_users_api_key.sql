-- *IMPORTANT: SQLC generates type-safe code from SQL
-- You write queries in SQL.
-- You run sqlc to generate code with type-safe interfaces to those queries.
-- You write application code that calls the generated code.

-- *IMPORTANT: Goose is a database migration tool. Both a CLI and a library.
-- It is used to create schemas for easier migrations and database management.
-- To do database migration Up, go to the sql/schema folder and use "goose postgres postgres://postgres:@localhost:5432/golang_rss_aggr up"
-- To do database migration Down, go to the sql/schema folder and use "goose postgres postgres://postgres:@localhost:5432/golang_rss_aggr down"

-- +goose Up
ALTER TABLE users ADD COLUMN api_key VARCHAR(64) UNIQUE NOT NULL DEFAULT (
    encode(sha256(random()::text::bytea), 'hex') -- creates a random unique value to store in api_key field
);

-- +goose Down
ALTER TABLE users DROP COLUMN api_key;