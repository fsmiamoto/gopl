package main

import (
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	reqParams := r.URL.Query()
	// Get cycle param from URL
	cycles, err := strconv.Atoi(reqParams.Get("cycle"))
	if err != nil {
		log.Fatal(err)
	}

	lissajous(w, &lParams{cycles: cycles})
}
