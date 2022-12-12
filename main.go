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

//	@title			Swagger Example API
//	@version		1.0
//	@description	This is a sample server celler server.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		localhost:8080
//	@BasePath	/api/v1

//	@securityDefinitions.basic	BasicAuth
func main() {
	var server *Server

	env := os.Getenv("ENV")
	switch env {
	case "DEV":
		server = newMinimalServer(DEV)
	case "UAT":
		server = newMinimalServer(UAT)
	case "PROD":
		server = newMinimalServer(PROD)
	default:
		server = newMinimalServer(DEV)
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
	case DEV:
		logger, _ = zap.NewDevelopment()
	case UAT, PROD:
		logger, _ = zap.NewProduction()
	}
	return logger.Sugar()
}

func newRouter(env Environment, l *zap.SugaredLogger) *chi.Mux {
	r := chi.NewRouter()

	r.Use(m.Logger(l))
	r.Use(middleware.Recoverer)

	return r
}
