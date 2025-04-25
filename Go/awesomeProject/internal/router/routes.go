package router

import (
	"github.com/faanross/orlokC2/internal/middleware"
	"github.com/go-chi/chi/v5"
)

func SetupRoutes(r chi.Router) {
	// Apply the middleware to all routes
	r.Use(middleware.UUIDMiddleware)

	// Register routes
	r.Get("/", RootHandler)
}
