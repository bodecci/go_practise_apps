package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utl-8")
	fmt.Fprint(w, "<h1>Welcome to my Awesome site!</h1>")
}

func pathHandler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		homeHandler(w, r)
	case "/contact":
		contactHanlder(w, r)
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "Page not found")
	}
	// fmt.Fprint(w, r.URL.Path)
}

func contactHanlder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utl-8")
	fmt.Fprint(w, "<h1>Contact Page</h1>")
}

func main() {
	http.HandleFunc("/", pathHandler)
	http.HandleFunc("/contact", contactHanlder)
	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", nil)
}
