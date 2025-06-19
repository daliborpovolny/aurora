package handlers

import (
	db "aurora/database"
	gen "aurora/database/gen"
	"database/sql"

	"aurora/internal"
	"aurora/internal/services"

	"context"
	"net/http"
)

type PublicHandler struct {
	Q   *gen.Queries
	Ctx context.Context
}

type PublicDeps struct {
	Q   *gen.Queries
	Ctx context.Context
}

func NewPublicHandler(f func(
	h PublicHandler,
	w http.ResponseWriter,
	r *http.Request)) internal.CustomHandler {

	return func(w http.ResponseWriter, r *http.Request) {
		h := PublicHandler{
			Q:   db.Queries,
			Ctx: r.Context(),
		}
		f(h, w, r)
	}
}

type PrivateDeps struct {
	Q       *gen.Queries
	Ctx     context.Context
	User    gen.User
	Session gen.Session
}

type PrivateHandler func(d PrivateDeps, w http.ResponseWriter, r *http.Request)

func NewPrivateHandler(f PrivateHandler) internal.CustomHandler {

	return func(w http.ResponseWriter, r *http.Request) {

		cookie, err := r.Cookie("session_cookie")
		if err != nil {
			if err == http.ErrNoCookie {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}
			http.Error(w, "Server error", http.StatusInternalServerError)
		}

		authInfo, err := services.UserService.GetAuthInfo(cookie.Value, r.Context())
		if err != nil {
			if err == sql.ErrNoRows {
				http.Error(w, "Invalid cookie", http.StatusForbidden)
			} else {
				panic(err)
			}
		}

		d := PrivateDeps{
			Q:       db.Queries,
			Ctx:     r.Context(),
			User:    authInfo.User,
			Session: authInfo.Session,
		}

		f(d, w, r)
	}
}
