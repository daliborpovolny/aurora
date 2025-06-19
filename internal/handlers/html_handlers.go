package handlers

import (
	db "aurora/database"
	"aurora/internal"
	"aurora/internal/services"
	"aurora/templates"
	"database/sql"
	"net/http"
)

type HtmlError struct {
	Message string `json:"message"`
}

type PublicHtmlHandler func(d PublicDeps, w http.ResponseWriter, r *http.Request) *HtmlError

func NewPublicHtmlHandler(f PublicHtmlHandler) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		d := PublicDeps{
			Q:   db.Queries,
			Ctx: r.Context(),
		}
		err := f(d, w, r)
		if err != nil {
			templates.ErrorBox(err.Message).Render(r.Context(), w)
		}
	}
}

type PrivateHtmlHandler func(d PrivateDeps, w http.ResponseWriter, r *http.Request) *HtmlError

func NewPrivateHtmlHandler(f PrivateHtmlHandler) internal.CustomHandler {

	return func(w http.ResponseWriter, r *http.Request) {

		cookie, err := r.Cookie("session_cookie")
		if err != nil {
			if err == http.ErrNoCookie {
				http.Redirect(w, r, "/login", http.StatusSeeOther)
				return
			}
			templates.ErrorBox(err.Error()).Render(r.Context(), w)
		}

		authInfo, err := services.UserService.GetAuthInfo(cookie.Value, r.Context())
		if err != nil {
			if err == sql.ErrNoRows {
				http.Redirect(w, r, "/login", http.StatusSeeOther)
			} else {
				templates.ErrorBox(err.Error()).Render(r.Context(), w)
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
