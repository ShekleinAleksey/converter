package main

import (
	"fmt"
	"net/http"
	"strconv"
)

type convertFunc func(float64) float64

var convertMap = map[string]map[string]convertFunc{
	"millimeters": {
		"centimeters": func(v float64) float64 { return v / 10 },
		"decimeters":  func(v float64) float64 { return v / 100 },
		"meters":      func(v float64) float64 { return v / 1000 },
	},
	"centimeters": {
		"millimeters": func(v float64) float64 { return v * 10 },
		"decimeters":  func(v float64) float64 { return v / 10 },
		"meters":      func(v float64) float64 { return v / 100 },
	},
	"decimeters": {
		"millimeters": func(v float64) float64 { return v * 100 },
		"centimeters": func(v float64) float64 { return v * 10 },
		"meters":      func(v float64) float64 { return v / 10 },
	},
	"meters": {
		"millimeters": func(v float64) float64 { return v * 1000 },
		"centimeters": func(v float64) float64 { return v * 100 },
		"decimeters":  func(v float64) float64 { return v * 10 },
	},
}

func main() {
	http.HandleFunc("/", convertHandler)

	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func convertHandler(w http.ResponseWriter, r *http.Request) {
	from := r.URL.Query().Get("from")
	value, err := strconv.ParseFloat(r.URL.Query().Get("value"), 64)
	if err != nil {
		http.Error(w, "Invalid value", http.StatusBadRequest)
		return
	}
	to := r.URL.Query().Get("to")
	convertFunc, ok := convertMap[from][to]
	if !ok {
		http.Error(w, "Invalid from or to unit", http.StatusBadRequest)
		return
	}

	result := convertFunc(value)

	fmt.Fprintf(w, "%f %s = %f %s", value, from, result, to)

}
