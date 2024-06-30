package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/google/uuid"
)

func getTimestamp() string {
	return time.Now().UTC().Format("2006-01-02T15:04:05Z")
}

func main() {
	uuid := uuid.New()

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT environment variable not set")
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		output := getTimestamp() + " " + uuid.String()
		fmt.Fprint(w, output)
	})

	log.Printf("Server started in port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))

	for {
		output := getTimestamp() + " " + uuid.String()
		fmt.Println(output)

		time.Sleep(time.Duration(5 * time.Second))
	}

}
