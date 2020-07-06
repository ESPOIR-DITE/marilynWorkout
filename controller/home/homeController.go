package home

import (
	"github.com/go-chi/chi"
	"html/template"
	"marilynWorkout/config"
	"net/http"
)

func HomeController(app *config.Env) http.Handler {
	r := chi.NewRouter()
	r.Get("/", HomeHandler(app))
	return r
}

func HomeHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		files := []string{
			app.Path + "index.html",
			app.Path + "template/header.html",
			//app.Path + "base_templates/footer.html",
		}
		ts, err := template.ParseFiles(files...)
		if err != nil {
			app.ErrorLog.Println(err.Error())
			return
		}
		err = ts.Execute(w, nil)
		if err != nil {
			app.ErrorLog.Println(err.Error())
		}
	}
}
