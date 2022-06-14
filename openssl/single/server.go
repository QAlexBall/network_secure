package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "this is https server page, single success!")
}

func main() {
	fmt.Println("Server Starting...")
	http.HandleFunc("/", handler)
	http.ListenAndServeTLS(":8082",
		"./server/server.crt", "./server/server.key", nil)
}
