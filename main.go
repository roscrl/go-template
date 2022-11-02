package main

import (
	"net/http"

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

type Server struct {
	env Environment
	log *zap.SugaredLogger

	router *chi.Mux
}

func newMinimalServer(env Environment) *Server {
	server := &Server{env: env, log: newLogger(env), router: newRouter(env)}
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

func newRouter(env Environment) *chi.Mux {
	router := chi.NewRouter()
	if env == DEV {
		router.Use(middleware.Logger)
	}

	router.Use(middleware.Recoverer)
	return router
}

func main() {
	server := newMinimalServer(DEV)
	defer server.log.Sync()

	err := http.ListenAndServe(":3000", server.router)
	if err != nil {
		server.log.Fatal(err)
	}
}
