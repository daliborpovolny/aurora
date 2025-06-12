package handlers

import (
	db "aurora/database"
	gen "aurora/database/gen"
	"aurora/internal"
	"context"
	"net/http"
)

type PublicHandler struct {
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
