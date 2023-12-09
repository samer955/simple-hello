package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1> Hello World 2 </h1>")
}

func main() {

	var port = "8080"
	http.HandleFunc("/", handler)
	log.Printf("Starting Server on Port %s\n", port)

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Println("Error: " + err.Error())
	}

}
