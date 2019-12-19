package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/cottrellio/cottrellio_go/pkg/creating"
	"github.com/cottrellio/cottrellio_go/pkg/db"
	"github.com/cottrellio/cottrellio_go/pkg/deleting"
	"github.com/cottrellio/cottrellio_go/pkg/reading"
	"github.com/cottrellio/cottrellio_go/pkg/routing"
	"github.com/cottrellio/cottrellio_go/pkg/updating"
)

const (
	appTimeout      = 10 * time.Second
	findWorkTimeout = 5 * time.Second
)

func main() {
	// Setup DB connection.
	db, err := db.New(db.MONGODB)
	if err != nil {
		log.Fatal(err)
	}

	// Setup Services.
	creator := creating.NewService(db)
	reader := reading.NewService(db)
	updater := updating.NewService(db)
	deleter := deleting.NewService(db)

	// Setup request handlers for routes.
	router := routing.Handler(creator, reader, updater, deleter)

	srv := &http.Server{
		Handler:      router,
		Addr:         ":8080",
		ReadTimeout:  appTimeout,
		WriteTimeout: appTimeout,
	}

	// Start Server
	go func() {
		log.Println("Starting server on port 8080...")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Error starting server: %s", err)
		}
	}()

	// Graceful Shutdown
	waitForShutdown(srv)
}

/* Shuts down the server without interrupting any active connections. First closes all open listeners,
   then closes all idle connections, and then waits till context timeout for connections to return to idle and then close them.
   If context expires before shutdown is complete, it returns context’s error, otherwise it returns any error returned
   from closing the server’s underlying listener.
*/
func waitForShutdown(srv *http.Server) {
	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// Wait until we receive shutdown signal.
	<-sigint

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		// Close database, truncate message queues, etc.
		cancel()
	}()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server shutdown failed:%+v", err)
	}

	log.Println("Server shutdown properly.")
	os.Exit(0)
}
