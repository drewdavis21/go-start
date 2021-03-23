// Package hello includes a test endpoint.
package hello

import (
	"github.com/drewdavis21/go-start/server"
	log "github.com/sirupsen/logrus"
	"net/http"
)

// init sets up the package for use.
func init() {
	addRoutes()
}

// HelloWorld writes back a nice greeting.
func HelloWorld(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("hello, world!"))
	if err != nil {
		log.Errorln(err)
	}
}

// addRoutes adds the routes for this package
func addRoutes() {
	server.BaseRouter.WithService("/", []server.Route{
		{
			Endpoint: "/",
			Handler:  HelloWorld,
			Method:   http.MethodGet,
		},
	})
}
