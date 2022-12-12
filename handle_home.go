package main

import (
	"net/http"
)

func (s *Server) handleHome() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		s.views.Home(w, req)
	}
}
