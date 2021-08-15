package main

import (
	"fmt"
	"net/http"
)

func headers(w http.ResponseWriter, r *http.Request) {
	h := r.Header
	fmt.Fprintln(w, h)
}

func body(w http.ResponseWriter, r *http.Request) {
	len := r.ContentLength
	body1 := make([]byte, len)

	r.Body.Read(body1)
	fmt.Fprintln(w, string(body1))

}

func main() {
	http.HandleFunc("/headers", headers)
	http.HandleFunc("/body", body)
	http.ListenAndServe(":8080", nil)
}
