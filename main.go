package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", dashboardHandler)
	http.HandleFunc("/data", dataHandler)
	http.HandleFunc("/table", tableHandler)
	http.HandleFunc("/events", sseHandler)
	http.Handle("/static", http.StripPrefix("/static", http.FileServer(http.Dir("static"))))
	http.ListenAndServe(":8080", nil)
}

func dashboardHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("dashboardHandler was called")
	tmpl := template.Must(template.ParseFiles("templates/dashboard.html"))
	fmt.Println("this was called 2nd")
	err := tmpl.Execute(w, nil)
	if err != nil {
		fmt.Println("Error executing template:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	fmt.Println("this was called 3rd")
}

func dataHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("dataHandler was called")
	tmpl := template.Must(template.ParseFiles("templates/data.html"))
	data := struct {
		Message string
	}{
		Message: "Jason sucks htmx!",
	}
	tmpl.Execute(w, data)
}

func tableHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("tableHandler was called")
	tmple := template.Must(template.ParseFiles("templates/table.html"))
	data := []struct {
		ID    int
		Name  string
		Value float64
	}{
		{1, "Test One", 123.45},
		{2, "Item Two", 678.98},
		{3, "Item Three", 456.71},
		{4, "Item Four", 456.73},
		{5, "Item Five", 56.44},
		{6, "Item Six", 456.70},
	}
	tmple.Execute(w, data)
}

func sseHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("sseHandler was called")
	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming unsupported!", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	for i := range 5 {
		fmt.Fprintf(w, "data: Message %d\n\n", i)
		flusher.Flush()
		time.Sleep(2 * time.Second)
	}
}
