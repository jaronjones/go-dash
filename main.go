package main

import (
	"html/template"
	"net/http"
)

func main() {
	http.HandleFunc("/", dashboardHandler)
	http.HandleFunc("/data", dataHandler)
	http.Handle("/static", http.StripPrefix("/static", http.FileServer(http.Dir("static"))))
	http.ListenAndServe(":8080", nil)
}

func dashboardHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/dashboard.html"))
	tmpl.Execute(w, nil)
}

func dataHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/data.html"))
	data := struct {
		Message string
	}{
		Message: "Jason sucks htmx!",
	}
	tmpl.Execute(w, data)
}
