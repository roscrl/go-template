package main

import (
	"app/models"
	"net/http"
)

func (s *Server) handlePosts() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		postsToDisplay := []models.Post{{Name: "aa", Author: "bbb"}, {Name: "aa", Author: "bbb"}}
		s.views.Posts(w, req, postsToDisplay)
	}
}
