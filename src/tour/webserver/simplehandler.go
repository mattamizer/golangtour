package main

import (
	"fmt"
	"log"
	"net/http"
)

type String string

type Struct struct {
	Greeting string
	Punct    string
	Who      string
}

func (s String) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, s)
}

func (s Struct) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, s)
}
func main() {
	// your http.Handle calls here
	http.Handle("/string", String("I'm a webserver!"))
	http.Handle("/struct", &Struct{"Hello", "!", "Dudeguybro"})
	log.Fatal(http.ListenAndServe("localhost:4000", nil))
}
