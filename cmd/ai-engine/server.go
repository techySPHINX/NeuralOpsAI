package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"neuralops/internal/domain"
	"neuralops/pkg/logging"
)

type Server struct {
	router *chi.Mux
	logger *logging.Logger
}

func NewServer(logger *logging.Logger) *Server {
	s := &Server{
		router: chi.NewRouter(),
		logger: logger,
	}

	s.router.Use(middleware.RequestID)
	s.router.Use(middleware.RealIP)
	s.router.Use(middleware.Logger)
	s.router.Use(middleware.Recoverer)

	s.routes()

	return s
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *Server) routes() {
	s.router.Get("/healthz", s.handleHealthz())

	s.router.Route("/v1", func(r chi.Router) {
		r.Post("/plan", s.handlePlan())
	})
}

func (s *Server) handleHealthz() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "OK")
	}
}

func (s *Server) handlePlan() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// In a real implementation, we would parse the NL query from the request
		// and use the AI engine to generate a plan.
		// For now, we return a dummy plan.

		plan := domain.PipelinePlan{
			ID:          "plan-123",
			Description: "A dummy pipeline plan",
			Tasks: []domain.Task{
				{
					Name:        "ingest-data",
					Description: "Ingest data from a source",
					Type:        "ingest",
					Config: map[string]string{
						"source": "s3://my-bucket/data.csv",
					},
				},
				{
					Name:        "transform-data",
					Description: "Transform the data",
					Type:        "transform",
					DependsOn:   []string{"ingest-data"},
					Config: map[string]string{
						"script": "SELECT * FROM input",
					},
				},
			},
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(plan)
	}
}
