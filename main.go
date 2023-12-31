package main

import (
	"fmt"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", handleHome)
	http.HandleFunc("/table", handleTable)
	http.HandleFunc("/stats", handleStats)

	loadEnv()

	if os.Getenv("ENV") == "dev" {
		fmt.Println("Dev server started and running at http://localhost:8080")
		log.Fatal(http.ListenAndServe("localhost:8080", nil))
	} else {
		fmt.Println("Server started and running")
		log.Fatal(http.ListenAndServe("0.0.0.0:"+os.Getenv("PORT"), nil))
	}
}

// handleHome is the handler for the home route ("/")
func handleHome(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.html", "templates/fragments.html"))
	err := tmpl.Execute(w, nil)
	if err != nil {
		log.Fatal(err)
	}
}

// handleStats is the handler for the stats route ("/stats")
func handleStats(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/stats.html", "templates/fragments.html"))
	err := tmpl.Execute(w, nil)
	if err != nil {
		log.Fatal(err)
	}
}

// handleTable is the handler for the table route ("/table"). The URL must
// include a query parameter "size" with a number to generate the table with
// the specified size. Otherwise, the handler will trigger a redirection back
// to the home route.
func handleTable(w http.ResponseWriter, r *http.Request) {
	sizeParam := r.URL.Query().Get("size")
	if sizeParam == "" {
		renderError(w, http.StatusBadRequest, "Query parameter `size` must be provided. It takes an integer between 3 and 6.")
		return
	}

	size, err := strconv.Atoi(sizeParam)

	if err != nil || size < 3 || size > 6 {
		renderError(w, http.StatusBadRequest, "Query parameter `size` must be an integer between 3 and 6.")
		return
	}

	nums := generateNums(size)

	showTimerParam := r.URL.Query().Get("timer")
	showTimer := showTimerParam == "y"

	tmpl := template.Must(
		template.ParseFiles(
			"templates/table.html",
			"templates/fragments.html",
		),
	)

	if err != nil {
		log.Fatal(err)
	}

	tmplData := struct {
		Nums      []int
		ShowTimer bool
	}{nums, showTimer}

	err = tmpl.Execute(w, tmplData)

	if err != nil {
		log.Fatal(err)
	}
}

// generateNums takes a table size (int) and generates a random sequence of
// numebrs necessary to generate the Schulte table
func generateNums(size int) []int {
	var nums []int

	for i := 1; i <= size*size; i++ {
		nums = append(nums, i)
	}

	for i := range nums {
		j := rand.Intn(size * size)
		nums[i], nums[j] = nums[j], nums[i]
	}

	return nums
}

// loadEnv simply loads .env file if it exists
func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		return
	}
}

// renderError renders an error template given a responseWriter, status code
// (int), and an error
// message (string)
func renderError(w http.ResponseWriter, errCode int, message string) {
	w.WriteHeader(errCode)

	tmpl := template.Must(
		template.ParseFiles(
			"templates/error.html",
			"templates/fragments.html",
		),
	)

	tmplData := struct {
		StatusCode int
		StatusText string
		Message    string
	}{errCode, http.StatusText(errCode), message}

	err := tmpl.Execute(w, tmplData)
	if err != nil {
		log.Fatal(err)
	}
}
