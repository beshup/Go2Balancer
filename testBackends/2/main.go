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
	fmt.Println("Serving on port 5000")
	http.ListenAndServe(":5000", nil)
}
