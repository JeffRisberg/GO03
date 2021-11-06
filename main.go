package main

import (
	"html/template"
	"net/http"
)

func charities(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Here are your charities"))
}

func donors(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Here are your donors"))
}

func process(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("tmpl.html")
	daysOfWeek := []string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}
	t.Execute(w, daysOfWeek)
}

func main() {
	http.HandleFunc("/charities", charities)
	http.HandleFunc("/donors", donors)
	http.HandleFunc("/process", process)

	http.Handle("/", http.FileServer(http.Dir("./src")))

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}

}
