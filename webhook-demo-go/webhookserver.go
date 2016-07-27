package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Receive request : %s %s\n", r.Method, r.URL.Path)
	switch r.Method {
	case "PUT":
		fmt.Fprintf(w, "Receive %s request\n", r.Method)
	case "GET":
		fmt.Fprintf(w, "Receive %s request\n", r.Method)
	case "POST":
		fmt.Fprintf(w, "Receive %s request\n", r.Method)
	default:
		fmt.Fprintf(w, "Receive unsupport method %s request\n", r.Method)
	}

	fmt.Fprintf(w, "Receive request : %s......", r.URL.Path[1:])
}

func main() {
	fmt.Println("\nHello webhook")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
