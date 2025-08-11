-- +goose Up
-- +goose StatementBegin
INSERT INTO user_table (username, name, sex, age, hash_pass, email)
VALUES ('karn', 'suksom', 'gay', 99, 'oasis', 'liam@noel.com');
INSERT INTO user_table (username, name, sex, age, hash_pass, email)
VALUES ('pla3', 'pla2', 'fish', 88, 'blur', 'song2@gmail.com');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM user_table;
-- +goose StatementEnd
