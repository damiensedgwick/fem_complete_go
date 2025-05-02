-- +goose Up
-- +goose StatementBegin
-- Create a new table with the desired structure
CREATE TABLE workouts_new (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    title TEXT NOT NULL,
    description TEXT,
    duration_minutes INTEGER NOT NULL,
    calories_burned INTEGER,
    created_at TEXT DEFAULT (datetime('now')),
    updated_at TEXT DEFAULT (datetime('now')),
    
    -- Add the new user_id column with the foreign key constraint
    user_id integer REFERENCES users(id) ON DELETE CASCADE
);

-- Copy all existing data to the new table with a default user_id value
INSERT INTO workouts_new 
SELECT *, 2
FROM workouts;

-- Drop the old table
DROP TABLE workouts;

-- Rename the new table to the original name
ALTER TABLE workouts_new RENAME TO workouts;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- Create a new table without the user_id column
CREATE TABLE workouts_new (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    title TEXT NOT NULL,
    description TEXT,
    duration_minutes INTEGER NOT NULL,
    calories_burned INTEGER,
    created_at TEXT DEFAULT (datetime('now')),
    updated_at TEXT DEFAULT (datetime('now')),
);

-- Copy data without the user_id column
INSERT INTO workouts_new
SELECT 
    id,
    title,
    description,
    duration_minutes,
    calories_burned,
    created_at,
    updated_at
FROM workouts;

-- Drop the old table
DROP TABLE workouts;

-- Rename the new table to the original name
ALTER TABLE workouts_new RENAME TO workouts;
-- +goose StatementEnd