package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")

		tpl := `
        <html>
        <body>
        <a href="/service-a">Service A</a>
        <br />
        <a href="/service-B">Service B</a>
        </body>

        </html>
        `
		fmt.Fprintf(w, tpl)
	})

	log.Println("Start Service Listen On :8080")

	http.ListenAndServe(":8080", nil)
}
