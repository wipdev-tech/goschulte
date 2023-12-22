package main

import (
	"fmt"
	"html/template"
	"log"
	"math"
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

	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

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

// handleTable is the handler for the table route ("/table"). The URL must
// include a query parameter "size" with a number to generate the table with
// the specified size. Otherwise, the handler will trigger a redirection back
// to the home route.
func handleTable(w http.ResponseWriter, r *http.Request) {
	sizeParam := r.URL.Query().Get("size")
	size, err := strconv.Atoi(sizeParam)
	if err != nil {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	nums := generateNums(size)

	showTimerParam := r.URL.Query().Get("timer")
	showTimer := showTimerParam == "y"

	tmpl, err := template.New("table.html").Funcs(
		template.FuncMap{"size": calcSize},
	).ParseFiles("templates/table.html", "templates/fragments.html")

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

// calcSize is a utility to be used in the table template. It takes an int
// slice and returns the square root of its length.
func calcSize(s []int) int {
	return int(math.Sqrt(float64(len(s))))
}
