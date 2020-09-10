package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Listening in port :3000")

	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})

	http.ListenAndServe(":3000", nil)
}