-- name: GetUser :one
SELECT * FROM user_table WHERE id = $1;

-- name: ListUsers :many
SELECT * FROM user_table;

-- name: GetUserbyEmail :one
SELECT * FROM user_table WHERE email = $1;

-- name: InsertUser :one
INSERT INTO user_table (username, name, sex, age, hash_pass, email)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING id, username, name, sex, age, hash_pass, email, image_path;

