package main

import (
	"basic-web/pkg/handlers"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println("Server started on :8080")
	http.ListenAndServe(":8080", nil)
}
