-- +goose Up
-- +goose StatementBegin
CREATE TABLE user_table (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    sex VARCHAR(255) NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE user_table;
-- +goose StatementEnd
