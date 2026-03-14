package main

import (
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		cyclesKey := "cycles"
		queries := r.URL.Query()

		if queries[cyclesKey] != nil {
			cycles, err := strconv.Atoi(queries[cyclesKey][0])
			if err != nil {
				log.Printf("Error parsing cycles query: %v", err)
				cycles = 5
			}
			lissajous(w, cycles)
			return
		}

		lissajous(w, 5)
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
