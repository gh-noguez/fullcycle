package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", Hello)
	err := http.ListenAndServe(":1078", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func Hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}
