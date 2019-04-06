package main

import (
	"net/http"

	"fmt"
)

func main() {

	http.HandleFunc("/zy", hello)

	http.ListenAndServe(":5050", nil)

}

func hello(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Hello Docker Form Golang!")

}
