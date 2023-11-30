package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", handleHome)
	http.HandleFunc("/table", handleTable)

	fmt.Println("Server started and running at http://localhost:8080")
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		tmpl := template.Must(template.ParseFiles("templates/index.html"))
		tmpl.Execute(w, nil)

	case "POST":
		size := r.FormValue("table-size")
		http.Redirect(w, r, "/table?size="+size, http.StatusFound)

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func handleTable(w http.ResponseWriter, r *http.Request) {
    sizeQueryParam := r.URL.Query().Get("size")
    size, err := strconv.Atoi(sizeQueryParam)
    if err != nil {
		http.Redirect(w, r, "/", http.StatusFound)
    }
    fmt.Println(size)

	tmpl := template.Must(template.ParseFiles("templates/table.html"))
	tmpl.Execute(w, nil)
}
