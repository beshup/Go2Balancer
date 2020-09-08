package main

import (
	"fmt"
	"net/http"
)

func test(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Whoa!")
}

func main() {
	http.HandleFunc("/", test)
	fmt.Println("Serving on port 4000")
	http.ListenAndServe(":4000", nil)
}
