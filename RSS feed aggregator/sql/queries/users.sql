-- The syntax for a sqlc query is like this->->name: {funcName} :{noOfRecordsToReturn}
-- After defining your schema use "sqlc generate" to generate the functions

-- name: CreateUser :one
INSERT INTO users (id, name, created_at, updated_at, api_key)
VALUES ($1, $2, $3, $4, 
    encode(sha256(random()::text::bytea), 'hex') -- creates a random unique value to store in api_key field
)
RETURNING *;

-- name: GetUserByAPIKey :one
SELECT * FROM users WHERE api_key = $1;