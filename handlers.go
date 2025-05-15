package main

import (
	database "aurora/database/gen"
	"context"
	"net/http"
)

type publicHandler struct {
	q   *database.Queries
	ctx context.Context
}

func newPublicHandler(f func(
	h publicHandler,
	w http.ResponseWriter,
	r *http.Request)) CustomHandler {

	return func(w http.ResponseWriter, r *http.Request) {
		h := publicHandler{
			q:   queries,
			ctx: r.Context(),
		}
		f(h, w, r)
	}
}
