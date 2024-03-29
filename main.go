package main

import (
  "api/config"
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
  e := NewRouter()

	// Configure server
	s := &http.Server{
		Addr:         "0.0.0.0:" + config.Port,
		ReadTimeout:  config.ReadTimeout * time.Second,
		WriteTimeout: config.WriteTimeout * time.Second,
		IdleTimeout:  config.IdleTimeout * time.Second,
	}

	// Start server
	go func() {
		if err := e.StartServer(s); err != nil {
			e.Logger.Info("Shutting down the server")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx := context.Background()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	} else {
		e.Logger.Info("Gracefully shutdown")
	}
}
