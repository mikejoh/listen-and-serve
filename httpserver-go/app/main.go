package main

import (
	"fmt"
	"log"
	"net/http"
)

func dumpRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Host: %q\n", r.Host)
	for k, v := range r.Header {
		fmt.Fprintf(w, "%q: %q\n", k, v)
	}
}

func main() {
	port := "5001"
	http.HandleFunc("/headers", dumpRequest)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
