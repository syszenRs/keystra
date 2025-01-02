package main

import (
	"keystone/api"
	"log"
)

const PORT = ":4000" //TODO: port as environment variable

func main() {
	log.Println("starting keystone project..")
	api.Start()
}
