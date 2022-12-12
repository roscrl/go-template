package main

import (
	"encoding/json"
	"net/http"
	"os"

	"app/views"

	m "app/middleware"

	_ "net/http/pprof"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"go.elastic.co/ecszap"
	"go.uber.org/zap"
)

type Environment int

const (
	LOCAL Environment = iota
	DEV
	UAT
	PROD
)

func (env Environment) String() string {
	return []string{"LOCAL", "DEV", "UAT", "PROD"}[env]
}

type Server struct {
	env Environment
	log *zap.SugaredLogger

	router *chi.Mux
	views  *views.View
}

// TODO fix formatting vim go templates
// TODO use go embed for templates

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
	var server *Server

	env := os.Getenv("ENV")
	switch env {
	case "LOCAL":
		server = newMinimalServer(LOCAL)
	case "DEV":
		server = newMinimalServer(DEV)
	case "UAT":
		server = newMinimalServer(UAT)
	case "PROD":
		server = newMinimalServer(PROD)
	default:
		server = newMinimalServer(LOCAL)
	}

	defer func() { _ = server.log.Sync() }()

	server.log.Infof("server started as %s on :3000", server.env)
	err := http.ListenAndServe(":3000", server.router)
	if err != nil {
		server.log.Fatal(err)
	}
}

func newMinimalServer(env Environment) *Server {
	log := newLogger(env)
	server := &Server{
		env:    env,
		log:    log,
		router: newRouter(env, log),
		views:  views.New(log),
	}

	server.routes()
	return server
}

func newLogger(env Environment) *zap.SugaredLogger {
	var logger *zap.Logger
	switch env {
	case LOCAL:
		logger, _ = zap.NewDevelopment()
	case DEV, UAT, PROD:
		config := zap.NewProductionConfig()
		config.EncoderConfig = ecszap.ECSCompatibleEncoderConfig(config.EncoderConfig)
		logger, _ = config.Build(ecszap.WrapCoreOption(), zap.AddCaller())
	}
	return logger.Sugar()
}

func newRouter(env Environment, l *zap.SugaredLogger) *chi.Mux {
	r := chi.NewRouter()

	r.Use(m.Logger(l))
	r.Use(middleware.Recoverer)

	return r
}
