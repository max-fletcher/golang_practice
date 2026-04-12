-- *IMPORTANT: SQLC generates type-safe code from SQL
-- You write queries in SQL.
-- You run sqlc to generate code with type-safe interfaces to those queries.
-- You write application code that calls the generated code.

-- *IMPORTANT: Goose is a database migration tool. Both a CLI and a library.
-- It is used to create schemas for easier migrations and database management.
-- To do database migration Up, use "goose postgres postgres://postgres:@localhost:5432/golang_rss_aggr up"
-- To do database migration Down, use "goose postgres postgres://postgres:@localhost:5432/golang_rss_aggr down"

-- +goose Up
CREATE TABLE users(
    id UUID PRIMARY KEY,
    name TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

-- +goose Down
DROP TABLE users;