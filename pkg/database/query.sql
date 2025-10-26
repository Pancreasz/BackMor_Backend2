-- name: GetUser :one
SELECT id, email, password_hash, display_name, avatar_url, bio, created_at, updated_at, sex, age
FROM users
WHERE id = $1;

-- name: ListUsers :many
SELECT id, email, password_hash, display_name, avatar_url, bio, created_at, updated_at, sex, age
FROM users;

-- name: GetUserbyEmail :one
SELECT * FROM users WHERE email = $1;

-- name: UpdateUserAvatarURL :one
UPDATE users
SET 
    avatar_url = $1,
    updated_at = CURRENT_TIMESTAMP
WHERE email = $2
RETURNING id, email, password_hash, display_name, avatar_url, bio, created_at, updated_at, sex, age;

-- name: InsertUser :one
INSERT INTO users (email, password_hash, display_name, avatar_url, bio, sex, age)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING id, email, password_hash, display_name, avatar_url, bio, created_at, updated_at, sex, age;

-- name: UpdateUserAvatarData :one
UPDATE users
SET 
    avatar_data = $1,
    updated_at = CURRENT_TIMESTAMP
WHERE email = $2
RETURNING id, email, password_hash, display_name, avatar_url, bio, created_at, updated_at, sex, age, avatar_data;



-- name: UpdateUserProfile :one
UPDATE users 
SET 
    display_name = COALESCE($1, display_name),
    avatar_url = COALESCE($2, avatar_url),
    bio = COALESCE($3, bio),
    sex = COALESCE($4, sex),
    age = COALESCE($5, age),
    updated_at = CURRENT_TIMESTAMP
WHERE email = $6
RETURNING id, email, password_hash, display_name, avatar_url, bio, created_at, updated_at, sex, age;