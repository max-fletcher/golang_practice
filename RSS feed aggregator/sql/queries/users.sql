-- The syntax for a sqlc query is like this->->name: {funcName} :{noOfRecordsToReturn}
-- After defining your schema use "sqlc generate" to generate the functions

-- name: CreateUser :one
INSERT INTO users (id, name, created_at, updated_at)
VALUES ($1, $2, $3, $4)
RETURNING *;