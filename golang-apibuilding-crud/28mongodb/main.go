package main

import (
	"fmt"
	"log"
	router "mongodb"
	"net/http"
)

func main() {
	fmt.Println("MongoDB API")
	r := router.Router() // not from mux
	fmt.Println("Server is getting started...")
	// http.ListenAndServe(":4000")
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("Listening at port 4000")

}
