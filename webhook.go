package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Webhook: %v\n\n\n", r)
	})

	port := ":80"
	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Println(err)
	}
}
