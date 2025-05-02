-- +goose Up
-- +goose StatementBegin
CREATE TABLE tokens_new (
    hash BLOB NOT NULL,
    user_id integer REFERENCES users(id) ON DELETE CASCADE,
    expires_at  created_at INTEGER,
    scope TEXT NOT NULL
);

-- Copy all existing data to the new table with a default user_id value
INSERT INTO tokens_new (hash, user_id, expires_at, scope)
SELECT hash, user_id, expires_at, scope FROM tokens;

-- Drop the old table
DROP TABLE tokens;

-- Rename the new table to the original name
ALTER TABLE tokens_new RENAME TO tokens;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- Create a new table without the user_id column
CREATE TABLE tokens_new (
    hash text PRIMARY KEY,
    user_id integer REFERENCES users(id) ON DELETE CASCADE,
    expires_at  created_at INTEGER,
    scope TEXT NOT NULL
);

-- Copy data without the changed column
INSERT INTO tokens_new (hash, user_id, expires_at, scope)
SELECT hash, user_id, expires_at, scope FROM tokens;

-- Drop the old table
DROP TABLE tokens;

-- Rename the new table to the original name
ALTER TABLE tokens_new RENAME TO tokens;
-- +goose StatementEnd