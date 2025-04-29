-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS workout_entries (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    workout_id INTEGER NOT NULL,
    exercise_name TEXT NOT NULL,
    sets INTEGER NOT NULL,
    reps INTEGER NOT NULL,
    duration_seconds INTEGER NOT NULL,
    weight REAL NOT NULL,
    notes TEXT,
    order_index INTEGER NOT NULL,
    created_at TEXT DEFAULT (datetime('now')),
    updated_at TEXT DEFAULT (datetime('now')),
    FOREIGN KEY (workout_id) REFERENCES workouts(id) ON DELETE CASCADE,
    CHECK (
        (reps IS NOT NULL OR duration_seconds IS NOT NULL) AND
        (reps IS NULL OR duration_seconds IS NULL)
    )
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS workout_entries;
-- +goose StatementEnd