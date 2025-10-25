-- +goose Up
-- +goose StatementBegin
ALTER TABLE users
ADD COLUMN avatar_data BYTEA;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE users
DROP COLUMN IF EXISTS avatar_data;
-- +goose StatementEnd
