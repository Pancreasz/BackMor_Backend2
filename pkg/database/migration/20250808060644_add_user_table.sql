-- +goose Up
-- +goose StatementBegin
CREATE TABLE user_table (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) UNIQUE NOT NULL,
    name VARCHAR(255) NOT NULL,
    sex VARCHAR(255) NOT NULL,
    age INT NOT NULL,
    hash_pass VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    image_path VARCHAR(500),
    created_timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE user_table;
-- +goose StatementEnd
