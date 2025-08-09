-- +goose Up
-- +goose StatementBegin
INSERT INTO user_table (id, name, sex)
VALUES (1, 'karn', 'gay');
INSERT INTO user_table (id, name, sex)
VALUES (2, 'pla3', 'fish');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM user_table;
-- +goose StatementEnd
