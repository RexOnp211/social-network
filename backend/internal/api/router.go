package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
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
	log.Printf("Incoming request: %s %s", req.Method, req.URL.Path) // リクエストログ

	path := req.URL.Path
	method := req.Method

	handler := r.getHandler(method, path)
	newHandler := corsHandler(handler)

	newHandler.ServeHTTP(w, req)
}

func (r *Router) getHandler(method, path string) http.Handler {
	for _, route := range r.routes {
		if route.Method == method && route.Pattern == path {
			return route.Handler
		} else if route.Pattern[len(route.Pattern)-1] == '/' && route.Pattern != "/" && strings.Contains(path, route.Pattern) && route.Method == method {
			return route.Handler
		}
	}
	return http.NotFoundHandler()
}

func corsHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		// You can customize the list of allowed origins
		allowedOrigins := map[string]bool{
			"http://localhost:3000": true,
			"http://app:3000":       true,
			"https://social-network-frontend-4ub2.onrender.com": true,
		}

		if allowedOrigins[origin] {
			w.Header().Set("Access-Control-Allow-Origin", origin)
		}
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
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
