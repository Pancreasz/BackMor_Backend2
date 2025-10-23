-- +goose Up
-- +goose StatementBegin
ALTER TABLE users ADD COLUMN sex VARCHAR(10);
ALTER TABLE users ADD COLUMN age INT;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE users DROP COLUMN sex;
ALTER TABLE users DROP COLUMN age;
-- +goose StatementEnd
