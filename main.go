package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	homeHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello, world!\n")
	}

	http.HandleFunc("/", homeHandler)
	log.Println("Listening for requsts at http://localhost:8000/")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
