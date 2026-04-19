-- *IMPORTANT: SQLC generates type-safe code from SQL
-- You write queries in SQL.
-- You run sqlc to generate code with type-safe interfaces to those queries.
-- You write application code that calls the generated code.

-- *IMPORTANT: Goose is a database migration tool. Both a CLI and a library.
-- It is used to create schemas for easier migrations and database management.
-- To do database migration Up, go to the sql/schema folder and use "goose postgres postgres://postgres:@localhost:5432/golang_rss_aggr up"
-- To do database migration Down, go to the sql/schema folder and use "goose postgres postgres://postgres:@localhost:5432/golang_rss_aggr down"

-- +goose Up
CREATE TABLE feed_follows(
    id UUID PRIMARY KEY,
    feed_id UUID NOT NULL REFERENCES feeds(id) ON DELETE CASCADE,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE, 
    UNIQUE(user_id, feed_id), -- a constraint that turns a combination of user_id and feed_id into a composite key
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

-- +goose Down
DROP TABLE feed_follows;