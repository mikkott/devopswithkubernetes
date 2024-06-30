package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func checkAndWriteFile(filePath string) {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		file, err := os.Create(filePath)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		_, err = file.WriteString("0")
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("File created and '0' written")
	} else {
		log.Printf("File already exists")
	}
}

func writeStringToFile(filePath string, content int) {
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	_, err = file.WriteString(fmt.Sprintf("%d", content))
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("String '%d' written to file", content)
}

func readIntFromFile(filePath string) (int, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	var content int
	_, err = fmt.Fscanf(file, "%d", &content)
	if err != nil {
		return 0, err
	}

	return content, nil
}

func main() {
	filepath := "/app/files/pingpong.txt"

	checkAndWriteFile(filepath)

	pingPongCount, err := readIntFromFile(filepath)

	if err != nil {
		log.Fatal(err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT environmental variable not set")
	}

	http.HandleFunc("/pingpong", func(w http.ResponseWriter, r *http.Request) {
		pingPongCount++
		writeStringToFile(filepath, pingPongCount)
		fmt.Fprint(w, "pong ", pingPongCount)
	})

	log.Printf("Server listening on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
