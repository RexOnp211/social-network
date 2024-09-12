package api

import (
	"encoding/json"
	"net/http"
	"regexp"
)

type Route struct {
	Method  string
	Pattern string
	Handler http.Handler
}

type Router struct {
	routes []Route
}

type Handler func(r *http.Request) (statusCode int, data map[string]interface{})

func NewRouter() *Router {
	return &Router{}
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	path := req.URL.Path
	method := req.Method

	handler := r.getHandler(method, path)
	newHandler := corsHandler(handler)

	newHandler.ServeHTTP(w, req)
}

func (r *Router) getHandler(method, path string) http.Handler {
	for _, route := range r.routes {
		re := regexp.MustCompile(route.Pattern)
		if route.Method == method && re.MatchString(path) {
			return route.Handler
		}
	}
	return http.NotFoundHandler()
}

func corsHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Access-Control-Allow-Methods", "Get, Post, Put, Delete")
		if r.Method == "OPTIONS" {
			return
		}
		next.ServeHTTP(w, r)
	})
}

func (r *Router) AddRoute(method, path string, handler http.Handler) {
	r.routes = append(r.routes, Route{Method: method, Pattern: path, Handler: handler})
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	statusCode, data := h(r)
	w.WriteHeader(statusCode)

	json.NewEncoder(w).Encode(data)
}

func (r *Router) GET(path string, handler Handler) {
	r.AddRoute("GET", path, handler)
}

func (r *Router) POST(path string, handler Handler) {
	r.AddRoute("POST", path, handler)
}

func (r *Router) PUT(path string, handler Handler) {
	r.AddRoute("PUT", path, handler)
}

func (r *Router) DELETE(path string, handler Handler) {
	r.AddRoute("DELETE", path, handler)
}
