package main

import (
	"fmt"
	"log"
	"net/http"

	// bringing third party libraries
	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Mod in golang")
	greeter()
	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/chat", servChat).Methods("GET") // for /chat API if we call the GET request, it will call the servChat function

	// running the above as a server 

	log.Fatal(http.ListenAndServe(":4000", r))
}

// go mod
// The job of go mod is to initialize it. It gives us the details on which version of go was the code written
// gorilla/mux ---> routing system used in GO
// Go toolchain is where we pool up all the dependencies for GO
// using go get -->  it fetches something directly from the web. Sometimes, our application is running on our server and we are not allowed to access the internet.In such cases we use (go get)
// go.sum --> prevents hacking --> it keeps track of what I am pulling from github. IT CONTAINS HASH AND IF ANY MALICIOUS THINGS HAPPEN THE HASH FAILS

func greeter() {
	fmt.Println("Hey there Mod users")
}

// for request and response
func serveHome(w http.ResponseWriter, r *http.Request) { // common syntax
	// r ---> somebody is sending us a request. We want to use the url and parameters which is present in r
	// w ---> contains the response

	w.Write([]byte("<h1>Welcome to golang series on youtube</h1>")) // send response back

}

func servChat(w http.ResponseWriter, r *http.Request) {
	fmt.Println("You have opened chat page")
	w.Write([]byte("<h1>Welcome to Chat Page</h1>"))
}
