package util

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// RunHTTPServerWithGracefulShutdown knows how to run and gracefully shutdown the HTTP Handler.
func RunHTTPServerWithGracefulShutdown(h http.Handler, addr string, shutdownTimeout time.Duration) error {
	log.Println("HTTP server is starting ...")

	// Create a server
	server := http.Server{
		Addr:    addr,
		Handler: h,
	}

	// Make a channel to listen for errors coming from the listener. Use a
	// buffered channel so the goroutine can exit if we don't collect this error.
	serverError := make(chan error, 1)

	// Start the service listening for requests.
	go func() {
		log.Printf("HTTP server is running on port: %v", server.Addr)
		serverError <- server.ListenAndServe()
	}()

	// Make a channel to listen for an interrupt or terminate signal from the OS.
	// Use a buffered channel because the signal package requires it.
	shutdownListenerChannel := make(chan os.Signal, 1)
	signal.Notify(shutdownListenerChannel, syscall.SIGINT, syscall.SIGTERM)

	// Blocking and waiting for shutdown or error from server.
	select {
	case err := <-serverError:
		if err != nil {
			log.Printf("Cannot start the HTTP server: %v", err)
			return err
		}
	case sig := <-shutdownListenerChannel:
		log.Printf("HTTP server shutdown by signal: %v", sig)

		// Give outstanding requests a deadline for completion.
		ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
		defer cancel()

		if err := server.Shutdown(ctx); err != nil {
			log.Printf("HTTP server was shutting down forcibly")
			err = server.Close()
			return err
		}

		log.Println("HTTP server was shutting down gracefully")
	}

	return nil
}
