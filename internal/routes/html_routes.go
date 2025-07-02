package routes

import (
	database "aurora/database/gen"
	"aurora/internal/handlers"
	"aurora/internal/services"
	"aurora/internal/utils"
	"aurora/templates"
	"errors"
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

func ViewStudentDetail(d handlers.PublicDeps, w http.ResponseWriter, r *http.Request) *handlers.HtmlError {
	student, err := services.StudentService.GetStudent(5, d.Ctx)
	if err != nil {
		return &handlers.HtmlError{
			Message: err.Error(),
		}
	}

	class, err := services.ClassService.GetClass(student.ClassID, d.Ctx)
	if err != nil {
		return &handlers.HtmlError{
			Message: err.Error(),
		}
	}

	teacher, err := services.TeacherService.GetTeacher(class.TeacherID, d.Ctx)
	if err != nil {
		return &handlers.HtmlError{
			Message: err.Error(),
		}
	}

	templates.StudentDetail(student, class, teacher).Render(d.Ctx, w)
	return nil
}

func Register(d handlers.PublicDeps, w http.ResponseWriter, r *http.Request) *handlers.HtmlError {

	var params services.RegisterParams
	err := utils.DecodeForm(r, &params)
	if err != nil {
		return &handlers.HtmlError{
			Message: "Bad form parameters",
		}
	}

	cookie, err := services.UserService.Register(params, d.Ctx)
	if err != nil {
		if errors.As(err, &services.EmailInUseErr) {
			return &handlers.HtmlError{
				Message: "Email is already in use by another account",
			}
		} else if errors.As(err, &services.BadPasswordErr) {
			return &handlers.HtmlError{
				Message: "Password is invalid because: " + err.Error(),
			}
		} else {
			return &handlers.HtmlError{
				Message: err.Error(),
			}
		}
	}

	r.AddCookie(cookie)
	http.Redirect(w, r, "/home", http.StatusSeeOther)
	return nil

}
