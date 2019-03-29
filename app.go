package main

import (
	"fmt"
	"log"
	"net/http"
)

func HomeEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello world :)")
}

func main() {
	http.HandleFunc("/", HomeEndpoint)
	if err := http.ListenAndServe(":80", nil); err != nil {
		log.Fatal(err)
	}
}
