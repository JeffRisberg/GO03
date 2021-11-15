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

type Data struct {
	Title   string
	Results []string
	Other   []int
}

func process2(w http.ResponseWriter, r *http.Request) {
	data := &Data{"Lord of the Rings", []string{"a", "b", "c"}, []int{1, 2, 3}}
	t, _ := template.ParseFiles("tmpl2.html")
	t.Execute(w, data)
}

type Donation struct {
	Charity string
	Amount  int
}

type DonationReport struct {
	Donations []Donation
}

func process3(w http.ResponseWriter, r *http.Request) {
	donation1 := &Donation{"Red Cross", 400}
	donation2 := &Donation{"Stanford University", 500}
	data := &DonationReport{[]Donation{*donation1, *donation2}}
	t, _ := template.ParseFiles("tmpl3.html")
	t.Execute(w, data)
}

func main() {
	http.HandleFunc("/charities", charities)
	http.HandleFunc("/donors", donors)
	http.HandleFunc("/process", process)
	http.HandleFunc("/process2", process2)
	http.HandleFunc("/process3", process3)

	http.Handle("/", http.FileServer(http.Dir("./src")))

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}

}
