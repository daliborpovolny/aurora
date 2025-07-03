package routes

import (
	database "aurora/database/gen"
	"aurora/internal/auth"
	"aurora/internal/handlers"
	"aurora/internal/services"
	"aurora/internal/utils"
	"aurora/templates"
	"errors"
	"fmt"
	"net/http"
)

//* html views

func ViewUsers(d handlers.PublicDeps, w http.ResponseWriter, r *http.Request) {
	users, err := services.UserService.ListUsers(d.Ctx)
	if err != nil {
		panic(err)
	}

	if users == nil {
		users = []database.User{}
	}

	cmp := templates.ListUsers(users, d.A)
	cmp.Render(d.Ctx, w)
}

func ViewStudents(d handlers.PublicDeps, w http.ResponseWriter, r *http.Request) {
	students, err := services.StudentService.ListStudents(d.Ctx)
	if err != nil {
		panic(err)
	}
	if students == nil {
		students = []database.ListStudentsRow{}
	}

	cmp := templates.ListStudents(students, d.A)
	cmp.Render(d.Ctx, w)
}

func ViewTeachers(d handlers.PublicDeps, w http.ResponseWriter, r *http.Request) {
	teachers, err := services.TeacherService.ListTeachers(d.Ctx)
	if err != nil {
		panic(err)
	}

	if teachers == nil {
		teachers = []database.ListTeachersRow{}
	}

	cmp := templates.ListTeachers(teachers, d.A)
	cmp.Render(d.Ctx, w)
}

func ViewParents(d handlers.PublicDeps, w http.ResponseWriter, r *http.Request) {
	parents, err := services.ParentService.ListParents(d.Ctx)
	if err != nil {
		panic(err)
	}

	if parents == nil {
		parents = []database.ListParentsRow{}
	}

	cmp := templates.ListParents(parents, d.A)
	cmp.Render(d.Ctx, w)
}

func ViewAdmins(d handlers.PublicDeps, w http.ResponseWriter, r *http.Request) {
	admins, err := services.AdminService.ListAdmins(d.Ctx)
	if err != nil {
		panic(err)
	}

	if admins == nil {
		admins = []database.ListAdminsRow{}
	}

	cmp := templates.ListAdmins(admins, d.A)
	cmp.Render(d.Ctx, w)
}

func ViewRegister(d handlers.PublicDeps, w http.ResponseWriter, r *http.Request) {
	cmp := templates.Register()
	cmp.Render(d.Ctx, w)
}

func ViewLogIn(d handlers.PublicDeps, w http.ResponseWriter, r *http.Request) {
	cmp := templates.Login()
	cmp.Render(d.Ctx, w)
}

func ViewStudentDetail(d handlers.PublicDeps, w http.ResponseWriter, r *http.Request) error {
	student, err := services.StudentService.GetStudent(5, d.Ctx)
	if err != nil {
		return err
	}

	class, err := services.ClassService.GetClass(student.ClassID, d.Ctx)
	if err != nil {
		return err
	}

	teacher, err := services.TeacherService.GetTeacher(class.TeacherID, d.Ctx)
	if err != nil {
		return err
	}

	templates.StudentDetail(student, class, teacher, d.A).Render(d.Ctx, w)
	return nil
}

func Register(d handlers.PublicDeps, w http.ResponseWriter, r *http.Request) error {
	var params auth.RegisterParams
	err := utils.DecodeForm(r, &params)
	if err != nil {
		return errors.New("unexpected form params")
	}

	cookie, err := auth.AuthService.Register(params, d.Ctx)
	if err != nil {
		if errors.As(err, &auth.EmailInUseError{}) {
			fmt.Println("Email in use")
			return err
		} else if errors.As(err, &auth.BadPasswordError{}) {
			fmt.Println("Bad password")
			return err
		} else {
			fmt.Println(err.Error())
			return err
		}
	}

	http.SetCookie(w, cookie)
	http.Redirect(w, r, "/", http.StatusSeeOther)
	return nil
}

func LogIn(d handlers.PublicDeps, w http.ResponseWriter, r *http.Request) error {

	var params auth.LoginParams
	err := utils.DecodeForm(r, &params)
	if err != nil {
		return errors.New("unexpected form params")
	}

	cookie, err := auth.AuthService.Login(params, d.Ctx)
	if err != nil {
		return err
	}

	http.SetCookie(w, cookie)
	http.Redirect(w, r, "/", http.StatusSeeOther)
	return nil
}

func Home(d handlers.PublicDeps, w http.ResponseWriter, r *http.Request) error {
	if r.URL.Path != "/" {
		return errors.New("unknown page")
	}

	cmp := templates.Home(d.A)
	cmp.Render(d.Ctx, w)
	return nil
}

func Count(d handlers.PublicDeps, w http.ResponseWriter, r *http.Request) error {
	fmt.Println(r.PathValue("id"))
	fmt.Fprint(w, "mmmmm :DD")
	return nil
}
