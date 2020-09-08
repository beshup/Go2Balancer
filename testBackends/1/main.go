package main

import (
	"fmt"
	"net/http"
)

func test(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Whoa, Go!")
}

func main() {
	http.HandleFunc("/", test)
	http.ListenAndServe(":4000", nil)
}
