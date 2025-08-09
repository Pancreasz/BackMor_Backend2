-- name: GetUser :one
SELECT * FROM user_table WHERE id = $1;

-- name: ListUsers :many
SELECT * FROM user_table;