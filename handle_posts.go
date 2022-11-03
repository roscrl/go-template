package main

import (
	"app/models"
	"app/views"
	"net/http"
)

func (s *Server) handlePosts() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		postsToDisplay := []models.Post{{Name: "Templ", Author: "John"}, {Name: "Chi", Author: "Bill"}, {Name: "Zap", Author: "Zip"}}
		views.Posts(postsToDisplay, w, req)
	}
}
