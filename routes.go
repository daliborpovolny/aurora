package main

import (
	"encoding/json"
	"net/http"

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
