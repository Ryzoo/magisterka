package main

import "net/http"

func main() {
	http.HandleFunc("/", testPage)
	http.ListenAndServe(":8080", nil)
}

func testPage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}