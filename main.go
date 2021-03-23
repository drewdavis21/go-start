package main

import (
	_ "github.com/drewdavis21/go-start/logging"
	"github.com/drewdavis21/go-start/server"
	_ "github.com/drewdavis21/go-start/hello"
)

// main begins execution of the server.
func main() {
	server.Serve()
}
