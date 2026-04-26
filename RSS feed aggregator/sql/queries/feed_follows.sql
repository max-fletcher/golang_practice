-- The syntax for a sqlc query is like this->->name: {funcName} :{noOfRecordsToReturn}
-- After defining your schema, go to your project root(where sqlc.yaml is) and use "sqlc generate" to generate the functions
-- For {noOfRecordsToReturn} when you are not returning a query result, you use :exec, but if you return one or many,
-- you use :one and :many respectively. :execrows is special because say for a delete operation, the query may delete multiple rows; 
-- it deletes and shows the number of records that were deleted

-- name: CreateFeedFollow :one
INSERT INTO feed_follows (id, feed_id, user_id, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetFeedFollowsByUserId :many
SELECT * 
FROM feed_follows
WHERE user_id=$1;

-- name: DeleteFeedFollowsById :execrows
DELETE FROM feed_follows 
WHERE id=$1 AND user_id=$2
RETURNING *;