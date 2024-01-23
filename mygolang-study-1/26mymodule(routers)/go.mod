module github/MeghanashreeC/mymodules

go 1.21.6

// indirect means this as of now it is not being used anywhere in the code. So as soon as you start using it, it disappears

require github.com/gorilla/mux v1.8.1


// go mod tidy
// go mod verify
// go list
// go list all 
// go list -m all
// go list -m -versions github.com/gorilla/mux
// go mod why github.com/gorilla/mux
// go mod graph
// go mod edit
// go mod edit -go 1.21.6
//go mod edit -module 1.21.6

// if we want to push all this to production

// go mod vendor
// go run -mod=vendor main.go
