package views

import (
	"app/models"
	"net/http"

	"github.com/a-h/templ"
)

func Home(w http.ResponseWriter, req *http.Request) {
	templ.Handler(home()).ServeHTTP(w, req)
}

func Posts(postsToDisplay []models.Post, w http.ResponseWriter, req *http.Request) {
	templ.Handler(posts(postsToDisplay)).ServeHTTP(w, req)
}
