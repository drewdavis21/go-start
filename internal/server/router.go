package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

// Router is a wrapper for the mux.Router type.
type Router struct {
	*mux.Router
}

// Route defines the structure for a route consisting of an
// endpoint and a handler function.
type Route struct {
	Endpoint string
	Handler  http.HandlerFunc
	Method   string
}

// BaseRouter is the base router for the server. It contains all of the routes.
var BaseRouter Router

// init sets up the package for use.
func init() {
	BaseRouter.Router = mux.NewRouter().StrictSlash(true)
	BaseRouter.Use(middleware)
}

// WithService adds a set of Routes to the given Router with the given path
// as a prefix.
func (r Router) WithService(path string, routes []Route) {
	sr := Router{r.Router.PathPrefix(path).Subrouter()}
	for _, r := range routes {
		sr.WithRoute(r)
	}
}

// WithRoute adds a Route to the given Router.
func (r Router) WithRoute(route Route) {
	r.HandleFunc(route.Endpoint, route.Handler).Methods(http.MethodOptions, route.Method)
}

// middleware adds logging to all routes on which it is used.
func middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Infoln(fmt.Sprintf("%s %s in %v", r.Method, r.URL.Path, time.Since(start)))
	})
}
