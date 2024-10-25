package main

import (
	"html/template"
	"net/http"
)

func main() {
	http.HandleFunc("/", dashboardHandler)
	http.HandleFunc("/data", dataHandler)
	http.HandleFunc("/table", tableHandler)
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

func tableHandler(w http.ResponseWriter, r *http.Request) {
	tmple := template.Must(template.ParseFiles("templates/table.html"))
	data := []struct {
		ID    int
		Name  string
		Value float64
	}{
		{1, "Item One", 123.45},
		{2, "Item Two", 678.98},
		{3, "Item Three", 456.71},
		{4, "Item Four", 456.73},
		{5, "Item Five", 56.44},
		{6, "Item Six", 456.70},
	}
	tmple.Execute(w, data)
}
