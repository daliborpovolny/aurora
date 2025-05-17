package main

import (
	"context"
	"database/sql"
	_ "embed"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	_ "modernc.org/sqlite"

	database "aurora/database/gen"
)

var resetDB bool = true
var port string = "8002"
var apiPrefix string = "/api/v1"

//go:embed schema.sql
var ddl string

var queries *database.Queries
var ctx context.Context

func home(w http.ResponseWriter, r *http.Request) {

	var n int
	as, err := queries.ListUsers(ctx)
	if err != nil {
		n = -1
	}
	n = len(as)

	fmt.Fprintf(w, "Hello, Worlds!, n: %d", n)
}

func getUsers(h publicHandler, w http.ResponseWriter, r *http.Request) {
	users, err := h.q.ListUsers(h.ctx)
	if err != nil {
		panic(err)
	}

	if users == nil {
		users = []database.User{}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(users)
	if err != nil {
		panic(err)
	}
}

func getStudents(h publicHandler, w http.ResponseWriter, r *http.Request) {
	students, err := h.q.ListStudents(h.ctx)
	if err != nil {
		panic(err)
	}
	if students == nil {
		students = []database.ListStudentsRow{}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(students)
	if err != nil {
		panic(err)
	}
}

func getTeachers(h publicHandler, w http.ResponseWriter, r *http.Request) {
	teachers, err := h.q.ListTeachers(h.ctx)
	if err != nil {
		panic(err)
	}

	if teachers == nil {
		teachers = []database.ListTeachersRow{}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(teachers)
	if err != nil {
		panic(err)
	}
}

func getParents(h publicHandler, w http.ResponseWriter, r *http.Request) {
	parents, err := h.q.ListParents(h.ctx)
	if err != nil {
		panic(err)
	}

	if parents == nil {
		parents = []database.ListParentsRow{}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(parents)
	if err != nil {
		panic(err)
	}
}

func getAdmins(h publicHandler, w http.ResponseWriter, r *http.Request) {
	admins, err := h.q.ListAdmins(h.ctx)
	if err != nil {
		panic(err)
	}

	if admins == nil {
		admins = []database.ListAdminsRow{}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(admins)
	if err != nil {
		panic(err)
	}
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

		queries.CreateParent(ctx, 1)
		queries.CreateTeacher(ctx, 2)
		queries.CreateStudent(ctx, 3)
		queries.CreateAdmin(ctx, 4)

		queries.AssignStudentToParent(ctx, database.AssignStudentToParentParams{
			ParentID:  1,
			StudentID: 1,
		})
	}

	r := NewRouter()

	r.GET("/home", home)

	r.GET(apiPrefix+"/users", newPublicHandler(getUsers))
	r.GET(apiPrefix+"/students", newPublicHandler(getStudents))
	r.GET(apiPrefix+"/teachers", newPublicHandler(getTeachers))
	r.GET(apiPrefix+"/admins", newPublicHandler(getAdmins))
	r.GET(apiPrefix+"/parents", newPublicHandler(getParents))

	s := &http.Server{
		Handler: r.ServeMux,
		Addr:    ":" + port,
	}

	fmt.Println("listing on ", s.Addr)

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
