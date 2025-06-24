package routes

import (
	database "aurora/database/gen"
	"aurora/internal/handlers"
	"aurora/internal/services"
	"aurora/templates"
	"net/http"
)

//* html views

func ViewUsers(h handlers.PublicHandler, w http.ResponseWriter, r *http.Request) {
	users, err := services.UserService.ListUsers(h.Ctx)
	if err != nil {
		panic(err)
	}

	if users == nil {
		users = []database.User{}
	}

	cmp := templates.ListUsers(users)
	cmp.Render(r.Context(), w)
}

func ViewStudents(h handlers.PublicHandler, w http.ResponseWriter, r *http.Request) {
	students, err := services.StudentService.ListStudents(h.Ctx)
	if err != nil {
		panic(err)
	}
	if students == nil {
		students = []database.ListStudentsRow{}
	}

	cmp := templates.ListStudents(students)
	cmp.Render(r.Context(), w)
}

func ViewTeachers(h handlers.PublicHandler, w http.ResponseWriter, r *http.Request) {
	teachers, err := services.TeacherService.ListTeachers(h.Ctx)
	if err != nil {
		panic(err)
	}

	if teachers == nil {
		teachers = []database.ListTeachersRow{}
	}

	cmp := templates.ListTeachers(teachers)
	cmp.Render(r.Context(), w)
}

func ViewParents(h handlers.PublicHandler, w http.ResponseWriter, r *http.Request) {
	parents, err := services.ParentService.ListParents(h.Ctx)
	if err != nil {
		panic(err)
	}

	if parents == nil {
		parents = []database.ListParentsRow{}
	}

	cmp := templates.ListParents(parents)
	cmp.Render(r.Context(), w)
}

func ViewAdmins(h handlers.PublicHandler, w http.ResponseWriter, r *http.Request) {
	admins, err := services.AdminService.ListAdmins(h.Ctx)
	if err != nil {
		panic(err)
	}

	if admins == nil {
		admins = []database.ListAdminsRow{}
	}

	cmp := templates.ListAdmins(admins)
	cmp.Render(r.Context(), w)
}

func ViewRegister(h handlers.PublicHandler, w http.ResponseWriter, r *http.Request) {
	cmp := templates.Register()
	cmp.Render(r.Context(), w)
}

func ViewLogIn(h handlers.PublicHandler, w http.ResponseWriter, r *http.Request) {
	cmp := templates.Login()
	cmp.Render(r.Context(), w)
}
