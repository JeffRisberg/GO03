package main

import (
	"net/http"
)

func charities(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Here are your charitices"))
}

func donors(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Here are your donors"))
}

func main() {
	http.HandleFunc("/charities", charities)
	http.HandleFunc("/donors", donors)

	http.Handle("/", http.FileServer(http.Dir("./src")))

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}

}
