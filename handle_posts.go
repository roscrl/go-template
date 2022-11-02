package main

import (
	"app/models"
	"app/views"
	"net/http"
)

func (s *Server) handlePosts() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		postsToDisplay := []models.Post{{Name: "aempl", Author: "author"}, {Name: "templ", Author: "author"}, {Name: "templ", Author: "author"}}
		views.Posts(postsToDisplay, w, req)
	}
}
