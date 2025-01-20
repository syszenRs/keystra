package main

import (
	"keystra/api"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	log.Println("starting keystra project..")
	var err error

	keystra, err := api.NewApi()

	if err != nil {
		log.Fatalf("Error creating API: %v", err)
	}

	go func() {
		if err = keystra.Start(); err != nil {
			log.Fatalf("Error starting server: %v", err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop

	log.Println("Shutting down...")

	if err = keystra.Close(); err != nil {
		log.Fatalf("Error closing API: %v", err)
	}
}
