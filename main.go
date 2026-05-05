package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "pong from cicd")
	})
	//没有注释
	http.ListenAndServe(":8080", nil)
}
