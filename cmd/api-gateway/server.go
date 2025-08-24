package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"neuralops/api/proto/ai_engine/v1"
	"neuralops/pkg/logging"
)

type Server struct {
	router    *chi.Mux
	logger    *logging.Logger
	aiClient  ai_enginev1.AIEngineServiceClient
}

func NewServer(logger *logging.Logger, aiClient ai_enginev1.AIEngineServiceClient) *Server {
	s := &Server{
		router:    chi.NewRouter(),
		logger:    logger,
		aiClient:  aiClient,
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
		r.Post("/pipelines:nl", s.handleCreatePipelineFromNL())
		r.Post("/pipelines/{id}:run", s.handleRunPipeline())
		r.Get("/runs/{id}", s.handleGetRun())
	})
}

func (s *Server) handleHealthz() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "OK")
	}
}

type NLQueryRequest struct {
	Query string `json:"query"`
}

func (s *Server) handleCreatePipelineFromNL() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req NLQueryRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		planResp, err := s.aiClient.Plan(context.Background(), &ai_enginev1.PlanRequest{Query: req.Query})
		if err != nil {
			s.logger.Error("failed to get plan from AI engine", "error", err)
			http.Error(w, "failed to create plan", http.StatusInternalServerError)
			return
		}

		s.logger.Info("Received plan from AI engine", "plan_id", planResp.Plan.Id)

		// In a future step, we will send this plan to the orchestrator.
		// For now, just return a success message.

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusAccepted)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "pipeline plan created successfully",
			"plan_id": planResp.Plan.Id,
		})
	}
}

func (s *Server) handleRunPipeline() http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        id := chi.URLParam(r, "id")
        w.WriteHeader(http.StatusAccepted)
        fmt.Fprintf(w, `{"message": "pipeline %s run started"}`)
    }
}

func (s *Server) handleGetRun() http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        id := chi.URLParam(r, "id")
        w.WriteHeader(http.StatusOK)
        fmt.Fprintf(w, `{"run_id": "%s", "status": "running"}`)
    }
}
