package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/configmap", CongifMap)
	http.HandleFunc("/", Hello)
	err := http.ListenAndServe(":1078", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
func CongifMap(w http.ResponseWriter, r *http.Request) {

	data, err := ioutil.ReadFile("/go/myfamily/family.txt")
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}
	fmt.Fprintf(w, "My family: %s", string(data))
}

func Hello(w http.ResponseWriter, r *http.Request) {

	name := os.Getenv("NAME")
	age := os.Getenv("AGE")

	fmt.Fprintf(w, "Hello %s, you are %s years old!", name, age)
}
