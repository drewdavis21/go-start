package main

import (
	_ "github.com/drewdavis21/go-start/internal/logging"
	"github.com/drewdavis21/go-start/internal/server"
	_ "github.com/drewdavis21/go-start/internal/test"
)

// main begins execution of the server.
func main() {
	server.Serve()
}
