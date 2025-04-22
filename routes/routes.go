package routes

import (
	"net/http"

	"github.com/dionofrizal88/go-allocator/config"
	"github.com/dionofrizal88/go-allocator/handler/webhook/agentallocation"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-redis/redis/v8"
)

// Router is a struct to save the router value.
type Router struct {
	config      config.Configuration
	redisClient *redis.Client
}

// NewRouter is a constructor that initializes a Router.
func NewRouter(options ...RouterOption) *Router {
	router := &Router{}

	for _, opt := range options {
		opt(router)
	}

	return router
}

// Init initializes and returns the chi router.
func (r *Router) Init() http.Handler {
	mux := chi.NewRouter()

	// Middleware stack
	mux.Use(middleware.Logger)
	mux.Use(middleware.Recoverer)

	// CORS
	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders: []string{"Link"},
		MaxAge:         300,
	}))

	// Health check or test API
	mux.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello this is chi!"))
	})

	agentAllocationController := agentallocation.NewController(r.config, r.redisClient)

	mux.Route("/api/v1/external/webhook", func(r chi.Router) {
		r.Post("/agent-allocation", func(w http.ResponseWriter, req *http.Request) {
			agentAllocationController.Manage(w, req)
		})
	})

	return mux
}
