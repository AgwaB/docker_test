package main

import "net/http"

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("Hello world"))
	})

	http.ListenAndServe(":8080", nil)
}