package api

import (
	"encoding/json"
	"fmt"
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
	path := req.URL.Path
	method := req.Method

	// process all requests with /profile/...
	//   if strings.HasPrefix(path, "/profile/") {
	//       corsHandler(http.HandlerFunc(handlers.ProfileHandler)).ServeHTTP(w, req)
	//       return
	//   }

	// // process all requests with /group/...
	// if strings.HasPrefix(path, "/group/") {
	//       corsHandler(http.HandlerFunc(handlers.GroupHandler)).ServeHTTP(w, req)
	//       return
	//   }

	// process other requests
	handler := r.getHandler(method, path)
	newHandler := corsHandler(handler)

	newHandler.ServeHTTP(w, req)
}

func (r *Router) getHandler(method, path string) http.Handler {
	for _, route := range r.routes {
		fmt.Println("Method", method, "Path", path, "Pattern", route.Pattern)
		if route.Method == method && route.Pattern == path {
			fmt.Println("Route mathed")
			return route.Handler
		} else if route.Pattern[len(route.Pattern)-1] == '/' && route.Pattern != "/" && strings.Contains(path, route.Pattern) && route.Method == method {
			fmt.Println("dynamic Route mathed")
			return route.Handler
		}
	}
	return http.NotFoundHandler()
}

func corsHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")

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
