package main

import (
	"fmt"
	"net/http"
)

func main() {
	// wip: todo: this is going to be where we try out ngrok-go

	port := ":80"
	fmt.Printf("Starting server on port %v\n\n", port)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Request: %v\n\n\n", r)

		fmt.Fprintf(w, "<h1>Authenticated User</h1>")
		fmt.Fprintf(w, "<div>Ngrok-Auth-User-Email: %v</div>", r.Header.Get("Ngrok-Auth-User-Email"))
		fmt.Fprintf(w, "<div>Ngrok-Auth-User-Id: %v</div>", r.Header.Get("Ngrok-Auth-User-Id"))
		fmt.Fprintf(w, "<div>Ngrok-Auth-User-Name: %v</div>", r.Header.Get("Ngrok-Auth-User-Name"))
		fmt.Fprintf(w, "<div>Referer: %v</div>", r.Header.Get("Referer"))
	})

	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Println(err)
	}
}
