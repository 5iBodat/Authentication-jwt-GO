package main

import (
	"authentication-jwt-go/src/router"
	"log"
	"net/http"
	"os"
)

func main() {
	log.Println("Server starting at http://localhost:%s", os.Getenv("PORT"))
	r := router.Router()
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
