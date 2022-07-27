package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("AAAAAAAAA"))
	})

	log.Println("Start Service A and Listen At: 8080")

	http.ListenAndServe(":8080", nil)
}
