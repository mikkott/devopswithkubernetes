package main

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

func getTimestamp() string {
	return time.Now().UTC().Format("2006-01-02T15:04:05Z")
}

func main() {
	uuid := uuid.New()

	for {
		output := getTimestamp() + " " + uuid.String()
		fmt.Println(output)

		time.Sleep(time.Duration(5 * time.Second))
	}

}
