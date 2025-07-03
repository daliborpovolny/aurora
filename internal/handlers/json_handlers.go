package handlers

import (
	db "aurora/database"
	"encoding/json"
	"net/http"
)

type JsonError interface {
	Status() int
	Message() string
	Error() string
}

type PublicJsonHandler func(d PublicDeps, w http.ResponseWriter, r *http.Request) JsonError

func NewPublicJsonHandler(f PublicJsonHandler) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		d := PublicDeps{
			Q:   db.Queries,
			Ctx: r.Context(),
		}
		err := f(d, w, r)
		if err != nil {

			errorMsg := struct {
				Status  int    `json:"status"`
				Message string `json:"message"`
			}{
				err.Status(),
				err.Error(),
			}

			json.NewEncoder(w).Encode(errorMsg)
		}
	}
}
