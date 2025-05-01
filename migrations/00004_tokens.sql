-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS tokens (
    hash text PRIMARY KEY,
    user_id integer REFERENCES users(id) ON DELETE CASCADE,
    expires_at  created_at INTEGER,
    scope TEXT NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS tokens;
-- +goose StatementEnd