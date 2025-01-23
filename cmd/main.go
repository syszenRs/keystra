package main

import (
	"keystra/api"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	args := os.Args
	log.Printf("Total args::%d:::%v", len(args), args)

	var apiStopFn func() error

	if len(args) < 2 {
		log.Println("API START")
		apiStopFn = startAPI()
	}

	log.Println("MORE")

	// Create a channel to listen for stop signals or manual "stop" commands
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	<-stop

	log.Println("Shutting down...")

	if err := apiStopFn(); err != nil {
		log.Fatalf("Error closing API: %v", err)
	}

}

func startAPI() func() error {
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

	return keystra.Close
}
