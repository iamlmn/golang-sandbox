package main

import ("fmt"
		"net/http")

func index_handler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Whoa, Go is Neat!")
}

func index_handler2(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Whoa, that was easily routed!")
}

func main() {
	http.HandleFunc("/", index_handler)
	http.HandleFunc("/about", index_handler2)
	http.ListenAndServe(":8000", nil)
}