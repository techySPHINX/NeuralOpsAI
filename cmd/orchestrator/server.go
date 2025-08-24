package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"neuralops/internal/domain"
	"neuralops/pkg/logging"
	"github.com/google/uuid"
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
		r.Post("/pipelines", s.handlePlanAndSubmit())
	})
}

func (s *Server) handleHealthz() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "OK")
	}
}

func (s *Server) handlePlanAndSubmit() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var plan domain.PipelinePlan
		if err := json.NewDecoder(r.Body).Decode(&plan); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// In a real implementation, we would compile the plan into a DAG
		// and submit it to Argo/Temporal.
		s.logger.Info("Received pipeline plan", "plan_id", plan.ID)

		runID := uuid.New().String()

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusAccepted)
		json.NewEncoder(w).Encode(map[string]string{"run_id": runID})
	}
}
