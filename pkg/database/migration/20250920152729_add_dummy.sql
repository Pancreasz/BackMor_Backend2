-- +goose Up
-- +goose StatementBegin
INSERT INTO users (email, password_hash, display_name)
SELECT
    'dummy' || g || '@example.com',
    'hashed_password_dummy',
    'Dummy User ' || g
FROM generate_series(1, 10) AS g;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM users
WHERE email LIKE 'dummy%@example.com';
-- +goose StatementEnd
