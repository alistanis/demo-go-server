package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		rw.WriteHeader(200)
		rw.Write([]byte("Hello!"))
	})

	mux.HandleFunc("/mark", func(rw http.ResponseWriter, r *http.Request) {
		rw.WriteHeader(200)
		rw.Write([]byte("Hello Mark!"))
	})

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
