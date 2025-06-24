package main

import (
	_ "embed"
	"fmt"
	"net/http"
	"os"

	_ "modernc.org/sqlite"

	db "aurora/database"
	"aurora/internal"
	"aurora/internal/handlers"
	"aurora/internal/routes"
	"aurora/templates"
)

var apiPrefix string = "/api/v1"

func home(w http.ResponseWriter, r *http.Request) {
	cmp := templates.Home()
	cmp.Render(r.Context(), w)
}

func main() {

	os.Setenv("RESET_DB", "true")
	os.Setenv("PORT", "8004")
	os.Setenv("IS_DEPLOYED", "false")

	db := db.Initialize()
	defer db.Close()

	r := internal.NewRouter()

	r.GET("/", home)

	r.GET(apiPrefix+"/users", handlers.NewPublicHandler(getUsers))
	r.GET(apiPrefix+"/students", handlers.NewPublicHandler(getStudents))
	r.GET(apiPrefix+"/teachers", handlers.NewPublicHandler(getTeachers))
	r.GET(apiPrefix+"/admins", handlers.NewPublicHandler(getAdmins))
	r.GET(apiPrefix+"/parents", handlers.NewPublicHandler(getParents))
	r.GET(apiPrefix+"/register", handlers.NewPublicHandler(apiRegister))

	r.GET("/users", handlers.NewPublicHandler(routes.ViewUsers))
	r.GET("/students", handlers.NewPublicHandler(routes.ViewStudents))
	r.GET("/teachers", handlers.NewPublicHandler(routes.ViewTeachers))
	r.GET("/parents", handlers.NewPublicHandler(routes.ViewParents))
	r.GET("/admins", handlers.NewPublicHandler(routes.ViewAdmins))

	r.GET("/register", handlers.NewPublicHandler(routes.ViewRegister))
	r.GET("/login", handlers.NewPublicHandler(routes.ViewLogIn))

	r.GET("/private", handlers.NewPrivateHtmlHandler(privateHome))

	s := &http.Server{
		Handler: r.ServeMux,
		Addr:    ":" + os.Getenv("PORT"),
	}

	fmt.Println("listening on ", s.Addr)

	err := s.ListenAndServe()
	if err != nil {
		fmt.Println(err)
	}
}
