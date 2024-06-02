package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("GET /hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World!!!")
	})

	http.ListenAndServe(":8080", nil)
}
