package main

import (
	"fmt"
	"net/http"
)

func process(w  http.ResponseWriter, r *http.Request)  {
	r.ParseMultipartForm(1024)
	//r.ParseForm()
	//map[hello:[sau sheong world] post:[456] thread:[123]]
	fmt.Fprintln(w, r.MultipartForm)

}

func main() {
	http.HandleFunc("/process", process)
	http.ListenAndServe(":8080", nil)
}
