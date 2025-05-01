package routes

import (
	"fem_complete_go/internal/app"

	"github.com/go-chi/chi/v5"
)

func SetupRoutes(app *app.Application) *chi.Mux {
	r := chi.NewRouter()

	r.Get("/health", app.HealthCheck)

	r.Get("/workouts/{id}", app.WorkoutHandler.HandleGetWorkingByID)
	r.Post("/workouts", app.WorkoutHandler.HandleCreateWorkout)
	r.Put("/workouts/{id}", app.WorkoutHandler.HandleUpdateWorkout)
	r.Delete("/workouts/{id}", app.WorkoutHandler.HandleDeleteWorkout)

	r.Post("/users", app.UserHandler.RegisterUser)

	return r
}
