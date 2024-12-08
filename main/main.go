package main

import (
	"cs/api"
	"cs/db"
	"net/http"
)

func main() {
	// Create a table
	db.CreateDb()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})
	http.HandleFunc("/login", api.Login)
	http.HandleFunc("/client", api.Clients)

	http.ListenAndServe(":8080", nil)
}
