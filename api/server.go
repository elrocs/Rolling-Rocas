package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.Method == http.MethodGet && r.URL.Path == "/ping":
		fmt.Fprintf(w, "pong")

	case r.Method == http.MethodGet && r.URL.Path == "/":
		http.ServeFile(w, r, "index.html")

	default:
		http.NotFound(w, r)
	}
}

func main() {
	http.HandleFunc("/", handler)

	fmt.Println("server listening at port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error al iniciar el servidor:", err)
	}
}
