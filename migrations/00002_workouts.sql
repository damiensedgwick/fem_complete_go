-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS workouts (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    -- user_id
    title TEXT NOT NULL,
    description TEXT,
    duration_minutes INTEGER NOT NULL,
    calories_burned INTEGER NOT NULL,
    created_at TEXT DEFAULT (datetime('now')),
    updated_at TEXT DEFAULT (datetime('now'))
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS workouts;
-- +goose StatementEnd