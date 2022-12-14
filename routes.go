package main

import (
	"github.com/go-chi/chi/v5/middleware"
)

func (s *Server) routes() {

	s.router.Get("/", s.handleHome())
	s.router.Get("/posts", s.handlePosts())

	s.router.Get("/swagger/*", s.handleSwagger())
	s.router.Mount("/debug", middleware.Profiler())
}
