package routes

import (
	"fem_complete_go/internal/app"

	"github.com/go-chi/chi/v5"
)

func SetupRoutes(app *app.Application) *chi.Mux {
	r := chi.NewRouter()

	r.Group(func(r chi.Router) {
		r.Use(app.Middleware.Authenticate)

		r.Post("/workouts", app.Middleware.RequireAuthenticatedUser(app.WorkoutHandler.HandleCreateWorkout))
		r.Get("/workouts/{id}", app.Middleware.RequireAuthenticatedUser(app.WorkoutHandler.HandleGetWorkingByID))
		r.Put("/workouts/{id}", app.Middleware.RequireAuthenticatedUser(app.WorkoutHandler.HandleUpdateWorkout))
		r.Delete("/workouts/{id}", app.Middleware.RequireAuthenticatedUser(app.WorkoutHandler.HandleDeleteWorkout))
	})

	r.Get("/health", app.HealthCheck)
	r.Post("/users", app.UserHandler.RegisterUser)
	r.Post("/tokens/authentication", app.TokenHandler.HandleCreateToken)

	return r
}
