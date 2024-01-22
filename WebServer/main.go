package main

import (
	"fmt"
	"net/http"
)

func main() {
	server := http.NewServeMux()
	server.HandleFunc("/hello", rootHandler)
	fileServer := http.FileServer(http.Dir("./public"))
	server.Handle("/", fileServer)
	err := http.ListenAndServe(":3333", server)
	if err != nil {
		fmt.Printf("\n Error Lunching web server : %v", err)
	}
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World! I am here to serve you"))
}
