package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo"
)

type Pongs struct {
	uuid  string
	pongs string
}

func getTimestamp() string {
	return time.Now().UTC().Format("2006-01-02T15:04:05Z")
}

func getHTTP(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func generateUUID() string {
	return uuid.New().String()
}

func main() {
	p := Pongs{uuid: generateUUID(), pongs: "0"}
	e := echo.New()
	e.GET("/", p.indexHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8081" // Default port if PORT environment variable is not set
	}

	e.Start(":" + port)

	for {
		output := getTimestamp() + " " + p.uuid
		fmt.Println(output)

		time.Sleep(time.Duration(5 * time.Second))
	}
}

func (p *Pongs) indexHandler(c echo.Context) error {
	timestamp := getTimestamp()
	pongs, err := getHTTP("http://pingpong:3456/pongs")
	if err != nil {
		return c.String(http.StatusInternalServerError, "Error getting pongs")
	}
	p.pongs = pongs
	return c.String(http.StatusOK, fmt.Sprintf(" %v: %v. \nPing / Pongs: %v", timestamp, p.uuid, p.pongs))
}
