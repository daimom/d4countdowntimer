package routes

import (
	"countdowntimer/controller"
	"net/http"

	"github.com/gorilla/mux"
)

var routes []Route

type Route struct {
	Method     string
	Pattern    string
	Handler    http.HandlerFunc
	Middleware mux.MiddlewareFunc
}

func init() {
	// ===== webhook =====
	register("GET", "/line/boss", controller.CheckTimer, nil)
}

func NewRouter() *mux.Router {
	r := mux.NewRouter()

	for _, route := range routes {
		r.Methods(route.Method, "OPTIONS").
			Path(route.Pattern).
			Handler(route.Handler)
	}

	return r
}

func register(method, pattern string, handler http.HandlerFunc, middleware mux.MiddlewareFunc) {
	routes = append(routes, Route{method, pattern, handler, middleware})
}
