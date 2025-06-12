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

	db.Initialize()

	r := internal.NewRouter()

	r.GET("/", home)

	r.GET(apiPrefix+"/users", handlers.NewPublicHandler(getUsers))
	r.GET(apiPrefix+"/students", handlers.NewPublicHandler(getStudents))
	r.GET(apiPrefix+"/teachers", handlers.NewPublicHandler(getTeachers))
	r.GET(apiPrefix+"/admins", handlers.NewPublicHandler(getAdmins))
	r.GET(apiPrefix+"/parents", handlers.NewPublicHandler(getParents))
	r.GET(apiPrefix+"/register", handlers.NewPublicHandler(apiRegister))

	r.GET("/users", handlers.NewPublicHandler(viewUsers))
	r.GET("/students", handlers.NewPublicHandler(viewStudents))
	r.GET("/teachers", handlers.NewPublicHandler(viewTeachers))
	r.GET("/parents", handlers.NewPublicHandler(viewParents))
	r.GET("/admins", handlers.NewPublicHandler(viewAdmins))

	r.GET("/register", handlers.NewPublicHandler(viewRegister))
	r.GET("/login", handlers.NewPublicHandler(viewLogIn))

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
