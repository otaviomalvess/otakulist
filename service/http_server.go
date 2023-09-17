package service

import (
	"context"
	"log"
	"net/http"
	"otavio-alves/OtakuList/configs"
	"otavio-alves/OtakuList/validation"
	"time"
)

// Creates the server
func createServer() (server *http.Server) {

	server = &http.Server{
		Addr:         configs.SERVER_ADDRS,
		ReadTimeout:  100 * time.Millisecond,
		WriteTimeout: 200 * time.Millisecond,
		IdleTimeout:  50 * time.Millisecond,
	}

	return
}

// StartServer .. Starts the server
func StartServer() {

	// Creates a validator
	validation.CreateValidator()

	// Creates a handler for the server
	handler := createHandler()

	// Creates an HTTP server
	server := createServer()

	// Creates a CORS struct to deal with cross origin
	cors := createCORS()

	// Associates a handler to the server
	server.Handler = cors.Handler(handler)

	// Makes sure to shutdown the server
	defer StopServer(server)

	// Starts the server and print error messages
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}

// StopServer .. Stops the server
func StopServer(server *http.Server) {

	// Creates a context for the operation
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)

	// Calls cancel context function
	defer cancel()

	// Shutdowns the server and checks an error occurs
	if err := server.Shutdown(ctx); err != nil {
		panic(err)
	}
}
