package controller

/**
* This is the main controller
* Every requests passe first here
* The purpose of this class is to direct the request(URL) coming from html to the respective controller classes
**/
import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"html/template"
	"marilynWorkout/config"
	"marilynWorkout/controller/home"
	"marilynWorkout/controller/plans/nutrition"
	"marilynWorkout/controller/plans/workout"
	"net/http"
)

func Controllers(env *config.Env) http.Handler {
	mux := chi.NewMux()
	mux.Use(middleware.RequestID)
	mux.Use(middleware.RealIP)
	mux.Use(middleware.Logger)
	mux.Use(env.Session.LoadAndSave)

	mux.Handle("/", HomeHandler(env))

	mux.Mount("/home", home.HomeController(env))
	mux.Mount("/plan_nutrition", nutrition.NutritionPlane(env))
	mux.Mount("/plan_workout", workout.WorkOutPlane(env))

	//mux.Mount("/book", book.Book(env))
	//mux.Mount("/category", item.Home(env))
	//mux.Mount("/customer", customer.Customer(env))
	//mux.Mount("/user", users.User(env))
	//mux.Mount("/manager", admin.Admin(env))
	//mux.Mount("/item", item.Home(env))
	//mux.Mount("/order", order.Order(env))

	fileServer := http.FileServer(http.Dir("./view/assets/"))
	// Use the mux.Handle() function to register the file server as the handler for
	// all URL paths that start with "/assets/". For matching paths, we strip the
	// "/static" prefix before the request reaches the file server.
	mux.Mount("/assets/", http.StripPrefix("/assets", fileServer))
	return mux
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
