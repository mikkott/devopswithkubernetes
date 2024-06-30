package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {

	pongCounter := -1

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT environment variable not set")
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		pongCounter++
		fmt.Fprint(w, "pong ", pongCounter)
	})

	log.Printf("Server started in port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))

}
