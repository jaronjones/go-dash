package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"time"

	log "github.com/sirupsen/logrus"
)

func main() {
	log.Info("Starting the server")
	http.HandleFunc("/", dashboardHandler)
	http.HandleFunc("/data", dataHandler)
	http.HandleFunc("/table", tableHandler)
	http.HandleFunc("/events", sseHandler)
	http.Handle("/static", http.StripPrefix("/static", http.FileServer(http.Dir("static"))))

	listener := http.ListenAndServe(":8080", nil)

	if listener != nil {
		log.Fatal("Failed to start the server")
	}
}

func init() {
	log.SetFormatter(&log.JSONFormatter{})

	log.SetOutput(os.Stdout)

	log.SetLevel(log.InfoLevel)
}

func dashboardHandler(w http.ResponseWriter, r *http.Request) {
	log.Info("dashboardHandler was called")

	tmpl := template.Must(template.ParseFiles("templates/dashboard.html"))

	err := tmpl.Execute(w, nil)
	if err != nil {
		fmt.Println("Error executing template:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func dataHandler(w http.ResponseWriter, r *http.Request) {
	log.Info("DataHandler was called")
	tmpl := template.Must(template.ParseFiles("templates/data.html"))
	data := struct {
		Message string
	}{
		Message: "Jason sucks htmx!",
	}

	err := tmpl.Execute(w, data)
	if err != nil {
		log.Error("Error executing data template")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func tableHandler(w http.ResponseWriter, r *http.Request) {
	log.Info("TableHandler was called")

	tmpl := template.Must(template.ParseFiles("templates/table.html"))

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

	err := tmpl.Execute(w, data)
	if err != nil {
		log.Error("TableHandler execution error")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func sseHandler(w http.ResponseWriter, r *http.Request) {
	log.Info("Executing event handler")
	flusher, ok := w.(http.Flusher)
	if !ok {
		log.Fatal("Client unable to support streaming!")
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
