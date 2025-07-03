package handlers

import (
	db "aurora/database"
	"aurora/internal"
	"aurora/internal/auth"
	"aurora/templates"
	"net/http"
)

type PublicHtmlHandler func(d PublicDeps, w http.ResponseWriter, r *http.Request) error

func NewPublicHtmlHandler(f PublicHtmlHandler) internal.CustomHandler {

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
		err = f(d, w, r)
		if err != nil {
			templates.ErrorBox(err.Error()).Render(r.Context(), w)
		}
	}
}

type PrivateHtmlHandler func(d PrivateDeps, w http.ResponseWriter, r *http.Request) error

func NewPrivateHtmlHandler(f PrivateHtmlHandler) internal.CustomHandler {

	return func(w http.ResponseWriter, r *http.Request) {

		authInfo, err := auth.AuthService.Authenticate(r)
		if err != nil {
			if err == http.ErrNoCookie {
				http.Redirect(w, r, "/login", http.StatusSeeOther)
			} else if err == auth.InvalidCookieErr {
				http.Redirect(w, r, "/login", http.StatusSeeOther)
			} else {
				templates.ErrorBox(err.Error()).Render(r.Context(), w)
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
