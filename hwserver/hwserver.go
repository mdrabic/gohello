package main

import (
	"log"
	"net/http"
	"io"
)

func main() {
	getHello := func(w http.ResponseWriter, _ *http.Request) {
		io.WriteString(w, "Hello!")
	}

	http.HandleFunc("/hello", getHello)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
