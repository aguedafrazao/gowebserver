package main

import (
	"fmt"
	"net/http"
)

func oi(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "oie")
}

func main() {
	http.HandleFunc("/oi", oi)
	http.ListenAndServe(":8080", nil)

}
