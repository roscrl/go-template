package views

import (
	"app/models"
	"net/http"
	"text/template"

	"go.uber.org/zap"
)

type View struct {
	Template *template.Template
	Log      *zap.SugaredLogger
}

func New(log *zap.SugaredLogger) *View {
	return &View{
		Template: template.Must(template.ParseGlob("views/*.html")),
		Log:      log,
	}
}

func (v *View) Home(w http.ResponseWriter, req *http.Request) {
	v.render(w, req, "home.html", nil)
}

func (v *View) Posts(w http.ResponseWriter, req *http.Request, postsToDisplay []models.Post) {
	v.render(w, req, "posts.html", postsToDisplay)
}

func (v *View) render(w http.ResponseWriter, req *http.Request, templateName string, data any) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	err := v.Template.Execute(w, data)
	if err != nil {
		http.Error(w, "Error rendering template for "+req.RequestURI, http.StatusInternalServerError)
		v.Log.Error(err)
	}
}
