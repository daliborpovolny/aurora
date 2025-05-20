package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	database "aurora/database/gen"
	templates "aurora/templates"
)

//* html views

func viewUsers(h publicHandler, w http.ResponseWriter, r *http.Request) {
	users, err := h.q.ListUsers(h.ctx)
	if err != nil {
		panic(err)
	}

	if users == nil {
		users = []database.User{}
	}

	cmp := templates.ListUsers(users)
	cmp.Render(r.Context(), w)
}
func viewStudents(h publicHandler, w http.ResponseWriter, r *http.Request) {
	students, err := h.q.ListStudents(h.ctx)
	if err != nil {
		panic(err)
	}
	if students == nil {
		students = []database.ListStudentsRow{}
	}

	cmp := templates.ListStudents(students)
	cmp.Render(r.Context(), w)
}

func viewTeachers(h publicHandler, w http.ResponseWriter, r *http.Request) {
	teachers, err := h.q.ListTeachers(h.ctx)
	if err != nil {
		panic(err)
	}

	if teachers == nil {
		teachers = []database.ListTeachersRow{}
	}

	cmp := templates.ListTeachers(teachers)
	cmp.Render(r.Context(), w)
}

func viewParents(h publicHandler, w http.ResponseWriter, r *http.Request) {
	parents, err := h.q.ListParents(h.ctx)
	if err != nil {
		panic(err)
	}

	if parents == nil {
		parents = []database.ListParentsRow{}
	}

	cmp := templates.ListParents(parents)
	cmp.Render(r.Context(), w)
}

func viewAdmins(h publicHandler, w http.ResponseWriter, r *http.Request) {
	admins, err := h.q.ListAdmins(h.ctx)
	if err != nil {
		panic(err)
	}

	if admins == nil {
		admins = []database.ListAdminsRow{}
	}

	cmp := templates.ListAdmins(admins)
	cmp.Render(r.Context(), w)
}

//* api endpoints

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

//* auth

type registerParams struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func register(h publicHandler, w http.ResponseWriter, r *http.Request) {

	var params registerParams
	err := Decode(r, &params)
	if err != nil {
		http.Error(w, "Failed to parse parameters", http.StatusBadRequest)
		return
	}

	if params.FirstName == "" || params.LastName == "" || params.Email == "" || params.Password == "" {
		http.Error(w, "Missing parameter", http.StatusBadRequest)
		return
	}

	hash, err := hashPassword(params.Password)
	if err != nil {
		fmt.Println("cannot hash")
		http.Error(w, "Unhashable password", http.StatusBadRequest)
		return
	}

	h.q.CreateUser(h.ctx, database.CreateUserParams{
		FirstName: params.FirstName,
		LastName:  params.LastName,
		Email:     params.Email,
		Hash:      hash,
	})

	value, err := NewSessionCookie()
	if err != nil {
		fmt.Println("cannot generate cookie")
		http.Error(w, "Server Error", http.StatusInternalServerError)
	}

	cookie := http.Cookie{
		Name:     "session_cookie",
		Value:    value,
		Path:     "/",
		HttpOnly: true,
		Secure:   isDeployed,
		Expires:  time.Now().Add(time.Hour * 7 * 24),
	}

	http.SetCookie(w, &cookie)
	fmt.Fprint(w, "Registered and set cookie")
}
