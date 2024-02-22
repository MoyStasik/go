package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Main page", r.URL)
}
func main() {
	http.HandleFunc("/", handler)

	http.HandleFunc("/page",
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintln(w, "Single page", r.URL)
		})
	http.HandleFunc("/pages/",
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, "Mutiple page", r.URL)
		})

	http.ListenAndServe(":8080", nil)

}
