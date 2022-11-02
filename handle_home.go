package main

import (
	"app/views"
	"net/http"
)

func (s *Server) handleHome() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		views.Home(w, req)
	}
}
