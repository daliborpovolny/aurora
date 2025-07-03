package handlers

import (
	db "aurora/database"
	gen "aurora/database/gen"

	"aurora/internal"
	"aurora/internal/auth"

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
	A   *auth.AuthInfo
}

func NewPublicHandler(f func(
	d PublicDeps,
	w http.ResponseWriter,
	r *http.Request)) internal.CustomHandler {

	return func(w http.ResponseWriter, r *http.Request) {

		authInfo, err := auth.AuthService.Authenticate(r)
		if err != nil {
			if err != http.ErrNoCookie && err != auth.InvalidCookieErr {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}

		d := PublicDeps{
			Q:   db.Queries,
			Ctx: r.Context(),
			A:   authInfo,
		}
		f(d, w, r)
	}
}

type PrivateDeps struct {
	Q   *gen.Queries
	Ctx context.Context
	A   *auth.AuthInfo
}

type PrivateHandler func(d PrivateDeps, w http.ResponseWriter, r *http.Request)

func NewPrivateHandler(f PrivateHandler) internal.CustomHandler {

	return func(w http.ResponseWriter, r *http.Request) {

		authInfo, err := auth.AuthService.Authenticate(r)
		if err != nil {
			if err == auth.InvalidCookieErr {
				http.Error(w, "invalid cookie, log in again", http.StatusUnauthorized)
			} else if err == http.ErrNoCookie {
				http.Error(w, "no cookie, log in", http.StatusUnauthorized)
			} else {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}

		d := PrivateDeps{
			Q:   db.Queries,
			Ctx: r.Context(),
			A:   authInfo,
		}

		f(d, w, r)
	}
}
