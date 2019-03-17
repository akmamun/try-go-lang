package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, Web!")
}
func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe(":8000", nil)
}

// func main() {
// 	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Fprintf(w, "Hello, Web %s", r.URL.Path)

// 	})
// 	http.ListenAndServe(":8080", nil)
// }
