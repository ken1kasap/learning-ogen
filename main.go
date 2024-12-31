package main

import (
	"log"
	"net/http"

	"github.com/ken1kasap/learning-ogen/handler"
	petstore "github.com/ken1kasap/learning-ogen/petstore"
)

func main() {
	// Create service instance.
	petsServiceHandler := handler.NewPetsService()

	// Create generated server.
	srv, err := petstore.NewServer(petsServiceHandler)
	if err != nil {
		log.Fatal(err)
	}
	if err := http.ListenAndServe(":8080", srv); err != nil {
		log.Fatal(err)
	}
}
