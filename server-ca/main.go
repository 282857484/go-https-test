package main

import (
	"fmt"
	"net/http"
)

func httpsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "this is an example http service")
}

func main() {
	http.HandleFunc("/", httpsHandler)
	http.ListenAndServeTLS(":8081",
		"server.crt",
		"server.key",nil)
}
