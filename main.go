package main

import (
	"context"
	"database/sql"
	_ "embed"
	"fmt"
	"net/http"
	"os"

	_ "modernc.org/sqlite"

	database "aurora/database/gen"
	"aurora/templates"
)

var resetDB bool = true
var port string = "8004"
var apiPrefix string = "/api/v1"
var isDeployed bool = false

//go:embed database/schema.sql
var ddl string

var queries *database.Queries
var ctx context.Context

func home(w http.ResponseWriter, r *http.Request) {

	cmp := templates.Home()
	cmp.Render(r.Context(), w)
}

func main() {
	if resetDB {
		os.Remove("./database/db_file.db")
	}

	ctx = context.Background()

	db, err := sql.Open("sqlite", "./database/db_file.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// create tables
	if _, err := db.ExecContext(ctx, ddl); err != nil {
		panic(err)
	}

	queries = database.New(db)

	if users, _ := queries.ListUsers(ctx); len(users) == 0 {
		queries.CreateUser(ctx, database.CreateUserParams{
			FirstName: "par",
			LastName:  "ent",
			Email:     "par.ent@gmail.com",
			Hash:      "parentHash",
		})

		queries.CreateUser(ctx, database.CreateUserParams{
			FirstName: "tea",
			LastName:  "cher",
			Email:     "tea.cher@isic.com",
			Hash:      "teachingHash",
		})

		queries.CreateUser(ctx, database.CreateUserParams{
			FirstName: "stu",
			LastName:  "end",
			Email:     "stu.dent@school.com",
			Hash:      "studentsHash",
		})

		queries.CreateUser(ctx, database.CreateUserParams{
			FirstName: "ad",
			LastName:  "min",
			Email:     "ad.min@manage.com",
			Hash:      "adminsHash",
		})

		_, err := queries.CreateParent(ctx, 1)
		if err != nil {
			panic(err)
		}

		teacher, err := queries.CreateTeacher(ctx, 2)
		if err != nil {
			panic(err)
		}

		class, err := queries.CreateClass(ctx, database.CreateClassParams{
			TeacherID:      teacher.ID,
			Room:           "A24",
			StartYear:      2024,
			GraduationYear: 2032,
		})
		if err != nil {
			panic(err)
		}

		queries.CreateStudent(ctx, database.CreateStudentParams{
			UserID:  3,
			ClassID: class.ID,
		})

		queries.CreateAdmin(ctx, 4)

		queries.AssignStudentToParent(ctx, database.AssignStudentToParentParams{
			ParentID:  1,
			StudentID: 1,
		})
	}

	r := NewRouter()

	r.GET("/", home)

	r.GET(apiPrefix+"/users", newPublicHandler(getUsers))
	r.GET(apiPrefix+"/students", newPublicHandler(getStudents))
	r.GET(apiPrefix+"/teachers", newPublicHandler(getTeachers))
	r.GET(apiPrefix+"/admins", newPublicHandler(getAdmins))
	r.GET(apiPrefix+"/parents", newPublicHandler(getParents))
	r.GET(apiPrefix+"/register", newPublicHandler(apiRegister))

	r.GET("/users", newPublicHandler(viewUsers))
	r.GET("/students", newPublicHandler(viewStudents))
	r.GET("/teachers", newPublicHandler(viewTeachers))
	r.GET("/parents", newPublicHandler(viewParents))
	r.GET("/admins", newPublicHandler(viewAdmins))

	r.GET("/register", newPublicHandler(viewRegister))
	r.GET("/login", newPublicHandler(viewLogIn))

	s := &http.Server{
		Handler: r.ServeMux,
		Addr:    ":" + port,
	}

	fmt.Println("listening on ", s.Addr)

	err = s.ListenAndServe()
	if err != nil {
		fmt.Println(err)
	}
}

// package main

// import (
// 	"fmt"
// 	"log"
// 	"net/http"

// 	"htmx/views"
// )

// type Count struct {
// 	Count int
// }

// func increment(w http.ResponseWriter, r *http.Request) {
// 	// Generate the component with a name
// 	component := views.Hello("Fucking hell this hurt")

// 	// Render the component using Templ's handler
// 	templ.Handler(component).ServeHTTP(w, r)
// }

// func main() {
// 	fmt.Println("Hello, world!")

// 	component := views.Hello("FUcking hell this hurt")
// 	http.Handle("/", templ.Handler(component))

// 	http.HandleFunc("/count", increment)

// 	log.Println("Server running on http://localhost:8080")
// 	http.ListenAndServe(":8080", nil)
// }
