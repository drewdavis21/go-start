package main

import (
	_ "github.com/drewdavis21/go-start/internal/logging"
	"github.com/drewdavis21/go-start/internal/server"
)

// main begins execution of the server.
func main() {
	server.Serve()
}
