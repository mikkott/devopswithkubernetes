package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/google/uuid"
)

func getTimestamp() string {
	return time.Now().UTC().Format("2006-01-02T15:04:05Z")
}

func readFileToString(filePath string) (string, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func calculateChecksum(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hash := sha256.New()
	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}

	checksum := hex.EncodeToString(hash.Sum(nil))
	return checksum, nil
}

func main() {
	filepath := "/app/files/output.log"

	uuid := uuid.New()

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT environment variable not set")
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		output, err := readFileToString(filepath)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		checksum, err := calculateChecksum(filepath)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Fprint(w, checksum, "\n", output)
	})

	log.Printf("Server started in port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))

	for {
		output := getTimestamp() + " " + uuid.String()
		fmt.Println(output)

		time.Sleep(time.Duration(5 * time.Second))
	}

}
