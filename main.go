package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "元神启动")
	})
	//没有注释
	http.ListenAndServe(":8080", nil)
}
