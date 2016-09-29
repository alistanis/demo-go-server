package main

import (
	"net/http"
)

func main(){
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request){
		rw.WriteHeader(200)
		rw.Write([]byte("Hello!")})
	})
}