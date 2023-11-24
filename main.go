package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func main() {
    fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
    http.HandleFunc("/", handleHome)

    fmt.Println("Server started and running at http://localhost:8080")
    log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func handleHome(w http.ResponseWriter, r *http.Request) {
    tmpl := template.Must(template.ParseFiles("templates/index.html"))
    tmpl.Execute(w, 42)
}
