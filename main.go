package main

import (
	"MTS/profile/httpserver"
	"log"
)

func main() {
	server := httpserver.New()

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
