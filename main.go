package main

import (
	"encoding/json"
	"net/http"

	m "app/middleware"

	_ "net/http/pprof"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"go.uber.org/zap"
)

type Environment int

const (
	DEV Environment = iota
	UAT
	PROD
)

func (env Environment) String() string {
	return []string{"DEV", "UAT", "PROD"}[env]
}

type Server struct {
	env Environment
	log *zap.SugaredLogger

	router *chi.Mux
}

func (s *Server) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	s.router.ServeHTTP(w, req)
}

func (s *Server) respond(w http.ResponseWriter, req *http.Request, data any, status int) {
	w.WriteHeader(status)
	if data == nil {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		s.log.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (s *Server) decode(w http.ResponseWriter, req *http.Request, data any) error {
	return json.NewDecoder(req.Body).Decode(data)
}

func main() {
	server := newMinimalServer(PROD)
	defer func() { _ = server.log.Sync() }()

	server.log.Infof("server started as %s on :3000", server.env)
	err := http.ListenAndServe(":3000", server.router)
	if err != nil {
		server.log.Fatal(err)
	}
}

func newMinimalServer(env Environment) *Server {
	log := newLogger(env)
	server := &Server{env: env, log: log, router: newRouter(env, log)}
	server.routes()
	return server
}

func newLogger(env Environment) *zap.SugaredLogger {
	var logger *zap.Logger
	switch env {
	case DEV:
		logger, _ = zap.NewDevelopment()
	case UAT, PROD:
		logger, _ = zap.NewProduction()
	}
	return logger.Sugar()
}

func newRouter(env Environment, l *zap.SugaredLogger) *chi.Mux {
	r := chi.NewRouter()

	if env == DEV {
		r.Use(middleware.Logger) // non structured chi dev logger
	} else {
		r.Use(m.Logger(l))
	}

	r.Use(middleware.Recoverer)
	return r
}
