-- name: GetUser :one
SELECT id, email, password_hash, display_name, avatar_url, bio, created_at, updated_at, sex, age
FROM users
WHERE id = $1;

-- name: ListUsers :many
SELECT id, email, password_hash, display_name, avatar_url, bio, created_at, updated_at, sex, age
FROM users;

-- name: GetUserbyEmail :one
SELECT * FROM users WHERE email = $1;

-- name: InsertUser :one
INSERT INTO users (email, password_hash, display_name, avatar_url, bio, sex, age)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING id, email, password_hash, display_name, avatar_url, bio, created_at, updated_at, sex, age;


-- name: GetUserByEmail :one
SELECT id, email, display_name, avatar_url, bio, created_at, updated_at, sex, age
FROM users
WHERE email = $1;