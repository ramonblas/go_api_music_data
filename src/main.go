package main

import (
	"log"
	"net/http"
	"go_api/src/controller"
)

func main() {
	router := controller.NewRouter()

	server := http.ListenAndServe(":8080", router)
	log.Fatal(server)
}
