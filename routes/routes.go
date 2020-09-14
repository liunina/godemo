/*
 * @Date: 2018-07-17 19:37:32
 * @LastEditors: liunian
 * @LastEditTime: 2020-09-14 17:28:42
 */
package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/liunina/godemo/auth"
	"github.com/liunina/godemo/controllers"
)

type Route struct {
	Method     string
	Pattern    string
	Handler    http.HandlerFunc
	Middleware mux.MiddlewareFunc
}

var routes []Route

func init() {
	register("GET", "/movies", controllers.AllMovies, auth.TokenMiddleware)
	register("GET", "/movies/{id}", controllers.FindMovie, nil)
	register("POST", "/movies", controllers.CreateMovie, nil)
	register("PUT", "/movies", controllers.UpdateMovie, nil)
	register("DELETE", "/movies/{id}", controllers.DeleteMovie, nil)

	register("POST", "/user/register", controllers.Register, nil)
	register("POST", "/user/login", controllers.Login, nil)
}

func NewRouter() *mux.Router {
	router := mux.NewRouter()
	for _, route := range routes {
		r := router.Methods(route.Method).
			Path(route.Pattern)
		if route.Middleware != nil {
			r.Handler(route.Middleware(route.Handler))
		} else {
			r.Handler(route.Handler)
		}
	}
	return router
}

func register(method, pattern string, handler http.HandlerFunc, middleware mux.MiddlewareFunc) {
	routes = append(routes, Route{method, pattern, handler, middleware})
}
