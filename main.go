package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	helloEnv := os.Getenv("HELLO")
	if helloEnv == "" {
		helloEnv = "Samer"
	}
	message := fmt.Sprintf("<h1> Hello %s </h1>", helloEnv)
	fmt.Fprintf(w, message)
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
