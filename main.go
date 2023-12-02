package main

import (
	"fmt"
	"html/template"
	"log"
	"math"
	"math/rand"
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

// handleHome is the handler for the home route ("/")
func handleHome(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.html", "templates/fragments.html"))
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
	sizeQueryParam := r.URL.Query().Get("size")
	size, err := strconv.Atoi(sizeQueryParam)
	if err != nil {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	nums := generateNums(size)

	tmpl, err := template.New("table.html").Funcs(
		template.FuncMap{"size": func(s []int) int {
			return int(math.Sqrt(float64(len(s))))
		}},
	).ParseFiles("templates/table.html", "templates/fragments.html")

	if err != nil {
		log.Fatal(err)
	}

	err = tmpl.Execute(w, nums)
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
