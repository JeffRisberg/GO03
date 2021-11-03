package main

import (
"fmt"
"net/http"
)

func main() {
	http.HandleFunc("/charities", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Here are your charities")
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
	})

	http.ListenAndServe(":80", nil)
}
