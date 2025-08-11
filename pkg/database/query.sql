-- name: GetUser :one
SELECT * FROM user_table WHERE id = $1;

-- name: ListUsers :many
SELECT * FROM user_table;

-- name: InsertUser :one
INSERT INTO user_table (name, sex)
VALUES ($1, $2)
RETURNING id, name, sex;

