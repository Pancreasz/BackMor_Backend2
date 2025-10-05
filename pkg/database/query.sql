-- name: GetUser :one
SELECT id, email, password_hash, display_name, avatar_url, bio, created_at, updated_at
FROM users
WHERE id = $1;

-- name: ListUsers :many
SELECT id, email, password_hash, display_name, avatar_url, bio, created_at, updated_at
FROM users;

-- name: GetUserbyEmail :one
SELECT * FROM user_table WHERE email = $1;

-- name: InsertUser :one
INSERT INTO users (email, password_hash, display_name, avatar_url, bio)
VALUES ($1, $2, $3, $4, $5)
RETURNING id, email, password_hash, display_name, avatar_url, bio, created_at, updated_at;


-- name: GetUserByEmail :one
SELECT id, email, display_name, avatar_url, bio, created_at, updated_at
FROM users
WHERE email = $1;