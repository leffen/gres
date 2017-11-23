package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/leffen/gres"

	"github.com/Sirupsen/logrus"
)

func main() {
	HTTPPort := 5017
	// Get the config

	// Initiate the logger
	logger := logrus.New()
	// logger.Formatter = &logrus.JSONFormatter{}

	r := gres.NewRouter(logger)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", HTTPPort),
		Handler: r,
	}

	go func() {
		err := srv.ListenAndServe()
		if err != nil {
			log.Fatal(err)
		}
	}()

	fmt.Printf("Listening on port :%d\n", conf.HTTPPort)

	// subscribe to SIGINT signals
	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, os.Interrupt)
	<-stopChan // wait for SIGINT
	log.Println("Shutting down server...")

	// shut down gracefully, but wait no longer than 5 seconds before halting
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	errS := srv.Shutdown(ctx)
	if errS != nil {
		log.Fatal(errS)
	}

	log.Println("Server gracefully stopped")
}
