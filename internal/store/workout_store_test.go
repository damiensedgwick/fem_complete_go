package store

import (
	"database/sql"
	"fmt"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupTestDB(t *testing.T) *sql.DB {
	err := godotenv.Load("../../.env")
	if err != nil {
		panic(err)
	}

	url := os.Getenv("TEST_DATABASE_URL")
	token := os.Getenv("TEST_AUTH_TOKEN")
	connectionString := fmt.Sprintf("%s?authToken=%s", url, token)

	db, err := sql.Open("libsql", connectionString)
	if err != nil {
		t.Fatalf("error opening database: %v", err)
	}

	err = Migrate(db, "../../migrations")
	if err != nil {
		t.Fatalf("error migrating database: %v", err)
	}

	_, err = db.Exec(`DELETE FROM workout_entries`)
	if err != nil {
		t.Fatalf("error clearing workout_entries table: %v", err)
	}

	_, err = db.Exec(`DELETE FROM workouts`)
	if err != nil {
		t.Fatalf("error clearing workouts table: %v", err)
	}

	_, err = db.Exec(`DELETE FROM sqlite_sequence WHERE name IN ('workouts', 'workout_entries')`)
	if err != nil {
		t.Fatalf("error resetting sequences: %v", err)
	}

	return db
}

func TestCreateWorkout(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	workoutStore := NewSqliteWorkoutStore(db)

	tests := []struct {
		name    string
		workout *Workout
		wantErr bool
	}{
		{
			name: "valid workout",
			workout: &Workout{
				Title:           "Push Day",
				Description:     "Upper Body Workouts",
				DurationMinutes: 60,
				CaloriesBurned:  300,
				Entries: []WorkoutEntry{
					{
						ID:           1,
						ExerciseName: "Bench Press",
						Sets:         3,
						Reps:         IntPtr(10),
						Weight:       IntFloatPtr(135),
						OrderIndex:   1,
						Notes:        "Good job!",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "workout with invalid entries",
			workout: &Workout{
				Title:           "full body",
				Description:     "complete body workout",
				DurationMinutes: 90,
				CaloriesBurned:  500,
				Entries: []WorkoutEntry{
					{
						ExerciseName: "Plank",
						Sets:         3,
						Reps:         IntPtr(60),
						Notes:        "keep form!",
						OrderIndex:   1,
					},
					{
						ExerciseName:    "Squats",
						Sets:            3,
						Reps:            IntPtr(12),
						DurationSeconds: IntPtr(60),
						Weight:          IntFloatPtr(135),
						Notes:           "full depth!",
						OrderIndex:      2,
					},
				},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			createdWorkout, err := workoutStore.CreateWorkout(tt.workout)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}

			require.NoError(t, err)
			assert.Equal(t, tt.workout.Title, createdWorkout.Title)
			assert.Equal(t, tt.workout.Description, createdWorkout.Description)
			assert.Equal(t, tt.workout.DurationMinutes, createdWorkout.DurationMinutes)
			// assert.Equal(t, tt.workout.CaloriesBurned, createdWorkout.CaloriesBurned)
			// assert.Equal(t, len(tt.workout.Entries), len(createdWorkout.Entries))
			// for i, entry := range tt.workout.Entries {
			// 	assert.Equal(t, entry.ExerciseName, createdWorkout.Entries[i].ExerciseName)
			// 	assert.Equal(t, entry.Sets, createdWorkout.Entries[i].Sets)
			// 	assert.Equal(t, entry.Reps, createdWorkout.Entries[i].Reps)
			// 	assert.Equal(t, entry.Weight, createdWorkout.Entries[i].Weight)
			// 	assert.Equal(t, entry.Notes, createdWorkout.Entries[i].Notes)
			// 	assert.Equal(t, entry.OrderIndex, createdWorkout.Entries[i].OrderIndex)
			// }

			retrievedWorkout, err := workoutStore.GetWorkoutByID(int64(createdWorkout.ID))
			require.NoError(t, err)
			assert.Equal(t, createdWorkout.ID, retrievedWorkout.ID)
			assert.Equal(t, len(tt.workout.Entries), len(retrievedWorkout.Entries))
			for i, entry := range tt.workout.Entries {
				assert.Equal(t, entry.ExerciseName, retrievedWorkout.Entries[i].ExerciseName)
				assert.Equal(t, entry.Sets, retrievedWorkout.Entries[i].Sets)
				assert.Equal(t, entry.OrderIndex, retrievedWorkout.Entries[i].OrderIndex)
			}
		})
	}
}

func IntPtr(i int) *int {
	return &i
}

func IntFloatPtr(f float64) *float64 {
	return &f
}
