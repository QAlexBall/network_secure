package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "this is https server page, single succes!")
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServeTLS(":8081",
		"../server.crt", "../server.key", nil)
}
