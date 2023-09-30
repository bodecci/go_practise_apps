package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utl-8")
	fmt.Fprint(w, "<h1>Welcome to my Awesome site!</h1>")
}

// func pathHandler(w http.ResponseWriter, r *http.Request) {
// 	switch r.URL.Path {
// 	case "/":
// 		homeHandler(w, r)
// 	case "/contact":
// 		contactHanlder(w, r)
// 	default:
// 		http.Error(w, "Page not found", http.StatusNotFound)
// 	}
// 	// fmt.Fprint(w, r.URL.Path)
// }

func contactHanlder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utl-8")
	fmt.Fprint(w, "<h1>Contact Page</h1>")
}

// have access to diff fields that we can set.
type Router struct{}

func (router Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		homeHandler(w, r)
	case "/contact":
		contactHanlder(w, r)
	default:
		http.Error(w, "Page not found", http.StatusNotFound)
	}
}

func main() {
	var router Router
	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", router)
}
