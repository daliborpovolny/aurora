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
)

var apiPrefix string = "/api/v1"

func main() {

	os.Setenv("RESET_DB", "false")
	os.Setenv("PORT", "8004")
	os.Setenv("IS_DEPLOYED", "false")

	db := db.Initialize()
	defer db.Close()

	r := internal.NewRouter()

	//TODO fix api, refactor to use json handler with jsonError
	// r.GET(apiPrefix+"/users", handlers.NewPublicHandler(getUsers))
	// r.GET(apiPrefix+"/students", handlers.NewPublicHandler(getStudents))
	// r.GET(apiPrefix+"/teachers", handlers.NewPublicHandler(getTeachers))
	// r.GET(apiPrefix+"/admins", handlers.NewPublicHandler(getAdmins))
	// r.GET(apiPrefix+"/parents", handlers.NewPublicHandler(getParents))
	// r.GET(apiPrefix+"/register", handlers.NewPublicHandler(apiRegister))

	r.GET("/", handlers.NewPublicHtmlHandler(routes.Home))

	r.GET("/users", handlers.NewPublicHandler(routes.ViewUsers))
	r.GET("/students", handlers.NewPublicHandler(routes.ViewStudents))
	r.GET("/teachers", handlers.NewPublicHandler(routes.ViewTeachers))
	r.GET("/parents", handlers.NewPublicHandler(routes.ViewParents))
	r.GET("/admins", handlers.NewPublicHandler(routes.ViewAdmins))

	r.GET("/register", handlers.NewPublicHandler(routes.ViewRegister))
	r.POST("/register", handlers.NewPublicHtmlHandler(routes.Register))

	r.GET("/login", handlers.NewPublicHandler(routes.ViewLogIn))
	r.POST("/login", handlers.NewPublicHtmlHandler(routes.LogIn))

	r.GET("/count/{id}", handlers.NewPublicHtmlHandler(routes.Count))

	// r.GET("/private", handlers.NewPrivateHtmlHandler(privateHome))

	s := &http.Server{
		Handler: r.ServeMux,
		Addr:    ":" + os.Getenv("PORT"),
	}

	r.ServeMux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	fmt.Println("listening on ", s.Addr)

	err := s.ListenAndServe()
	if err != nil {
		fmt.Println(err)
	}
}
