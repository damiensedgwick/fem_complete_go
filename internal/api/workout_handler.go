package api

import (
	"encoding/json"
	"fem_complete_go/internal/store"
	"fem_complete_go/internal/utils"
	"log"
	"net/http"
)

type WorkoutHandler struct {
	workoutStore store.WorkoutStore
	logger       *log.Logger
}

func NewWorkoutHandler(workoutStore store.WorkoutStore, logger *log.Logger) *WorkoutHandler {
	return &WorkoutHandler{
		workoutStore: workoutStore,
		logger:       logger,
	}
}

func (wh *WorkoutHandler) HandleGetWorkingByID(w http.ResponseWriter, r *http.Request) {
	workoutID, err := utils.ReadIDParam(r)
	if err != nil {
		wh.logger.Printf("ERROR: ReadIDParam %v", err)
		utils.WriteJSON(w, http.StatusBadRequest, utils.Envelope{"error": "invalid workout id"})
		return
	}

	workout, err := wh.workoutStore.GetWorkoutByID(workoutID)
	if err != nil {
		wh.logger.Printf("ERROR: GetWorkoutByID %v", err)
		utils.WriteJSON(w, http.StatusInternalServerError, utils.Envelope{"error": "internal server error"})
		return
	}

	utils.WriteJSON(w, http.StatusOK, utils.Envelope{"workout": workout})
}

func (wh *WorkoutHandler) HandleCreateWorkout(w http.ResponseWriter, r *http.Request) {
	var workout store.Workout
	if err := json.NewDecoder(r.Body).Decode(&workout); err != nil {
		wh.logger.Printf("ERROR: DecodingCreateWorkout%v", err)
		utils.WriteJSON(w, http.StatusBadRequest, utils.Envelope{"error": "invalid request body"})
		return
	}

	createdWorkout, err := wh.workoutStore.CreateWorkout(&workout)
	if err != nil {
		wh.logger.Printf("ERROR: CreateWorkout%v", err)
		utils.WriteJSON(w, http.StatusInternalServerError, utils.Envelope{"error": "failed to create workout"})
		return
	}

	utils.WriteJSON(w, http.StatusCreated, utils.Envelope{"workout": createdWorkout})
}

func (wh *WorkoutHandler) HandleUpdateWorkout(w http.ResponseWriter, r *http.Request) {
	workoutID, err := utils.ReadIDParam(r)
	if err != nil {
		wh.logger.Printf("ERROR: ReadIDParam %v", err)
		utils.WriteJSON(w, http.StatusBadRequest, utils.Envelope{"error": "invalid workout update id"})
		return
	}

	existingWorkout, err := wh.workoutStore.GetWorkoutByID(workoutID)
	if err != nil {
		wh.logger.Printf("ERROR: GetWorkoutByID %v", err)
		utils.WriteJSON(w, http.StatusInternalServerError, utils.Envelope{"error": "internal server error"})
		return
	}

	if existingWorkout == nil {
		http.NotFound(w, r)
		return
	}

	var updateWorkoutRequest struct {
		Title           *string              `json:"title"`
		Description     *string              `json:"description"`
		DurationMinutes *int                 `json:"duration_minutes"`
		CaloriesBurned  *int                 `json:"calories_burned"`
		Entries         []store.WorkoutEntry `json:"entries"`
	}

	if err := json.NewDecoder(r.Body).Decode(&updateWorkoutRequest); err != nil {
		wh.logger.Printf("ERROR: DecodeUpdateRequest %v", err)
		utils.WriteJSON(w, http.StatusBadRequest, utils.Envelope{"error": "invalid request payload"})
		return
	}

	if updateWorkoutRequest.Title != nil {
		existingWorkout.Title = *updateWorkoutRequest.Title
	}

	if updateWorkoutRequest.Description != nil {
		existingWorkout.Description = *updateWorkoutRequest.Description
	}

	if updateWorkoutRequest.DurationMinutes != nil {
		existingWorkout.DurationMinutes = *updateWorkoutRequest.DurationMinutes
	}

	if updateWorkoutRequest.CaloriesBurned != nil {
		existingWorkout.CaloriesBurned = *updateWorkoutRequest.CaloriesBurned
	}

	if updateWorkoutRequest.Entries != nil {
		existingWorkout.Entries = updateWorkoutRequest.Entries
	}

	if err := wh.workoutStore.UpdateWorkout(existingWorkout); err != nil {
		wh.logger.Printf("ERROR: UpdateWorkout %v", err)
		utils.WriteJSON(w, http.StatusInternalServerError, utils.Envelope{"error": "failed to update workout"})
		return
	}

	utils.WriteJSON(w, http.StatusOK, utils.Envelope{"workout": existingWorkout})
}

func (wh *WorkoutHandler) HandleDeleteWorkout(w http.ResponseWriter, r *http.Request) {
	workoutID, err := utils.ReadIDParam(r)
	if err != nil {
		wh.logger.Printf("ERROR: ReadIDParam %v", err)
		utils.WriteJSON(w, http.StatusBadRequest, utils.Envelope{"error": "invalid workout update id"})
		return
	}

	if err := wh.workoutStore.DeleteWorkout(workoutID); err != nil {
		wh.logger.Printf("ERROR: DeleteWorkout %v", err)
		utils.WriteJSON(w, http.StatusInternalServerError, utils.Envelope{"error": "failed to delete workout"})
		return
	}

	utils.WriteJSON(w, http.StatusOK, utils.Envelope{"message": "workout deleted successfully"})
}
