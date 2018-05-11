package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func server(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	// send response to the client
	fmt.Fprintf(w, "the time is %v", time.Now())
}

func main() {
	http.HandleFunc("/", server)             // set router
	err := http.ListenAndServe(":9090", nil) // set listen port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
