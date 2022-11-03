package main

import (
	"encoding/json"
	"net/http"
)

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
