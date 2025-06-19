package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	database "aurora/database/gen"
	"aurora/templates"

	"aurora/internal/auth"
	"aurora/internal/handlers"
	"aurora/internal/services"
	"aurora/internal/utils"
)

//* html views

func viewUsers(h handlers.PublicHandler, w http.ResponseWriter, r *http.Request) {
	users, err := h.Q.ListUsers(h.Ctx)
	if err != nil {
		panic(err)
	}

	if users == nil {
		users = []database.User{}
	}

	cmp := templates.ListUsers(users)
	cmp.Render(r.Context(), w)
}
func viewStudents(h handlers.PublicHandler, w http.ResponseWriter, r *http.Request) {
	students, err := h.Q.ListStudents(h.Ctx)
	if err != nil {
		panic(err)
	}
	if students == nil {
		students = []database.ListStudentsRow{}
	}

	cmp := templates.ListStudents(students)
	cmp.Render(r.Context(), w)
}

func viewTeachers(h handlers.PublicHandler, w http.ResponseWriter, r *http.Request) {
	teachers, err := h.Q.ListTeachers(h.Ctx)
	if err != nil {
		panic(err)
	}

	if teachers == nil {
		teachers = []database.ListTeachersRow{}
	}

	cmp := templates.ListTeachers(teachers)
	cmp.Render(r.Context(), w)
}

func viewParents(h handlers.PublicHandler, w http.ResponseWriter, r *http.Request) {
	parents, err := h.Q.ListParents(h.Ctx)
	if err != nil {
		panic(err)
	}

	if parents == nil {
		parents = []database.ListParentsRow{}
	}

	cmp := templates.ListParents(parents)
	cmp.Render(r.Context(), w)
}

func viewAdmins(h handlers.PublicHandler, w http.ResponseWriter, r *http.Request) {
	admins, err := h.Q.ListAdmins(h.Ctx)
	if err != nil {
		panic(err)
	}

	if admins == nil {
		admins = []database.ListAdminsRow{}
	}

	cmp := templates.ListAdmins(admins)
	cmp.Render(r.Context(), w)
}

func viewRegister(h handlers.PublicHandler, w http.ResponseWriter, r *http.Request) {
	cmp := templates.Register()
	cmp.Render(r.Context(), w)
}

func viewLogIn(h handlers.PublicHandler, w http.ResponseWriter, r *http.Request) {
	cmp := templates.Login()
	cmp.Render(r.Context(), w)
}

//* api endpoints

func getUsers(h handlers.PublicHandler, w http.ResponseWriter, r *http.Request) {
	users, err := h.Q.ListUsers(h.Ctx)
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

func getStudents(h handlers.PublicHandler, w http.ResponseWriter, r *http.Request) {
	students, err := h.Q.ListStudents(h.Ctx)
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

func getTeachers(h handlers.PublicHandler, w http.ResponseWriter, r *http.Request) {
	teachers, err := h.Q.ListTeachers(h.Ctx)
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

func getParents(h handlers.PublicHandler, w http.ResponseWriter, r *http.Request) {
	parents, err := h.Q.ListParents(h.Ctx)
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

func getAdmins(h handlers.PublicHandler, w http.ResponseWriter, r *http.Request) {
	admins, err := h.Q.ListAdmins(h.Ctx)
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

//* auth

type registerParams struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func apiRegister(h handlers.PublicHandler, w http.ResponseWriter, r *http.Request) {

	var params registerParams
	err := utils.DecodeJson(r, &params)
	if err != nil {
		http.Error(w, "Failed to parse parameters", http.StatusBadRequest)
		return
	}

	if params.FirstName == "" || params.LastName == "" || params.Email == "" || params.Password == "" {
		http.Error(w, "Missing parameter", http.StatusBadRequest)
		return
	}

	hash, err := auth.HashPassword(params.Password)
	if err != nil {
		fmt.Println("cannot hash")
		http.Error(w, "Unhashable password", http.StatusBadRequest)
		return
	}

	h.Q.CreateUser(h.Ctx, database.CreateUserParams{
		FirstName: params.FirstName,
		LastName:  params.LastName,
		Email:     params.Email,
		Hash:      hash,
	})

	value, err := auth.NewSessionCookie()
	if err != nil {
		fmt.Println("cannot generate cookie")
		http.Error(w, "Server Error", http.StatusInternalServerError)
	}

	cookie := http.Cookie{
		Name:     "session_cookie",
		Value:    value,
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
		Expires:  time.Now().Add(time.Hour * 7 * 24),
	}

	http.SetCookie(w, &cookie)

	json.NewEncoder(w).Encode(utils.JSONResponse{
		Status:  http.StatusCreated,
		Message: "Registered succesfully",
	})
}

func privateHome(d handlers.PrivateDeps, w http.ResponseWriter, r *http.Request) *handlers.HtmlError {
	fmt.Fprint(w, "Email: "+d.User.Email)
	return nil
}

func htmlRegister(d handlers.PublicDeps, w http.ResponseWriter, r *http.Request) *handlers.HtmlError {

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
