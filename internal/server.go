package internal

import (
	"net/http"
)

type CustomHandler func(w http.ResponseWriter, r *http.Request)

func (h CustomHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h(w, r)
}

type Router struct {
	ServeMux *http.ServeMux
	handlers map[string]map[string]http.Handler
}

func NewRouter() *Router {
	return &Router{
		ServeMux: http.NewServeMux(),
		handlers: make(map[string]map[string]http.Handler),
	}
}

func (r *Router) Handle(method string, pattern string, handler http.Handler) {
	if _, exists := r.handlers[pattern]; !exists {
		r.handlers[pattern] = make(map[string]http.Handler)
	}
	r.handlers[pattern][method] = handler

	r.ServeMux.HandleFunc(pattern, func(w http.ResponseWriter, req *http.Request) {
		if handler, ok := r.handlers[pattern][req.Method]; ok {
			handler.ServeHTTP(w, req)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
}

// func (r *Router) ServeHTTP()

// GET registers a new GET request handler with the given pattern
func (r *Router) GET(pattern string, handler CustomHandler) {
	r.Handle("GET", pattern, handler)
}

// POST registers a new POST request handler with the given pattern
func (r *Router) POST(pattern string, handler CustomHandler) {
	r.Handle("POST", pattern, handler)
}

// PUT registers a new PUT request handler with the given pattern
func (r *Router) PUT(pattern string, handler CustomHandler) {
	r.Handle("PUT", pattern, handler)
}

// DELETE registers a new DELETE request handler with the given pattern
func (r *Router) DELETE(pattern string, handler CustomHandler) {
	r.Handle("DELETE", pattern, handler)
}
