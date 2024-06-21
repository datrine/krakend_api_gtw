-- +goose Up
-- +goose StatementBegin
CREATE TRIGGER IF NOT EXISTS users_updated_at_trigger AFTER UPDATE ON users BEGIN UPDATE users SET updated_at = datetime(); END
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TRIGGER IF EXISTS users_updated_at_trigger;
-- +goose StatementEnd
