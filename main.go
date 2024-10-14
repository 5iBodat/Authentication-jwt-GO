package main

import (
	"authentication-jwt-go/src/router"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	fmt.Printf("Server starting at http://localhost:%s", os.Getenv("PORT"))
	r := router.Router()
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
