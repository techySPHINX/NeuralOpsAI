package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
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

func (s *Server) handleCreatePipelineFromNL() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusAccepted)
		fmt.Fprintln(w, `{"message": "pipeline creation from natural language started"}`)
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
