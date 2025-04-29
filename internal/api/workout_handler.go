package api

import (
	"encoding/json"
	"fem_complete_go/internal/store"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type WorkoutHandler struct {
	workoutStore store.WorkoutStore
}

func NewWorkoutHandler(workoutStore store.WorkoutStore) *WorkoutHandler {
	return &WorkoutHandler{
		workoutStore: workoutStore,
	}
}

func (wh *WorkoutHandler) HandleGetWorkingByID(w http.ResponseWriter, r *http.Request) {
	paramsWorkoutID := chi.URLParam(r, "id")
	if paramsWorkoutID == "" {
		http.NotFound(w, r)
		return
	}

	workoutID, err := strconv.ParseInt(paramsWorkoutID, 10, 64)
	if err != nil {
		http.NotFound(w, r) // this is a cheat for now, we'll fix this later
		return
	}

	fmt.Fprintf(w, "this is the workout by id: %d\n", workoutID)
}

func (wh *WorkoutHandler) HandleCreateWorkout(w http.ResponseWriter, r *http.Request) {
	var workout store.Workout
	if err := json.NewDecoder(r.Body).Decode(&workout); err != nil {
		fmt.Printf("decode error: %s\n", err) // this is just during development
		http.Error(w, "failed to create workout", http.StatusInternalServerError)
		return
	}

	createdWorkout, err := wh.workoutStore.CreateWorkout(&workout)
	if err != nil {
		fmt.Printf("create error: %s\n", err) // this is just during development
		http.Error(w, "failed to create workout", http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(createdWorkout)
}
