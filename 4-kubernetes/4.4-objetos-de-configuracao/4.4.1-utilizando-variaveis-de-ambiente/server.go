package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", Hello)
	err := http.ListenAndServe(":1078", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func Hello(w http.ResponseWriter, r *http.Request) {

	name := os.Getenv("NAME")
	age := os.Getenv("AGE")

	fmt.Fprintf(w, "Hello %s, you are %s years old!", name, age)
}
