package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo"
)

type Pongs struct {
	pongs int
}

func main() {
	e := echo.New()
	p := Pongs{pongs: 0}

	e.GET("/pingpong", p.pingPongHandler)
	e.GET("/pongs", p.pongsHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8081" // Default port if PORT environment variable is not set
	}
	e.Start(":" + port)
}

func (p *Pongs) pingPongHandler(c echo.Context) error {
	p.pongs++
	return c.String(http.StatusOK, fmt.Sprintf("pong %v", p.pongs))
}

func (p *Pongs) pongsHandler(c echo.Context) error {
	return c.String(http.StatusOK, fmt.Sprint(p.pongs))
}
