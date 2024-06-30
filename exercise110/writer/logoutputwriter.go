package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/google/uuid"
)

func getTimestamp() string {
	return time.Now().UTC().Format("2006-01-02T15:04:05Z")
}

func writeStringToFile(filename string, content string) {
	f, err := os.OpenFile(filename, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	if _, err := f.WriteString(content + "\n"); err != nil {
		log.Fatal(err)
	}
}

func main() {
	filepath := "/app/files/output.log"

	uuid := uuid.New()

	for {
		output := getTimestamp() + " " + uuid.String()
		fmt.Println(output)

		writeStringToFile(filepath, output)

		time.Sleep(time.Duration(5 * time.Second))
	}

}
