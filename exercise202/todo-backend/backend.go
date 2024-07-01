package main

import (
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo"
)

type Todo struct {
	Title string `json:"title"`
}

type Todos struct {
	Todos []Todo `json:"todos"`
}

func main() {
	e := echo.New()
	t := Todos{
		Todos: []Todo{
			{Title: "Todo 1"},
			{Title: "Todo 2"},
			{Title: "Todo 3"},
		},
	}
	e.GET("/todos", t.getTodosHandler)
	e.POST("/todos", t.postTodosHandler)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081" // Default port if PORT environment variable is not set
	}
	e.Start(":" + port)
}

func (t *Todos) getTodosHandler(c echo.Context) error {
	log.Println("Sending todos:", t.Todos)
	return c.JSON(http.StatusOK, t.Todos)
}

func (t *Todos) postTodosHandler(c echo.Context) error {
	title := c.FormValue("todo")
	log.Println("Adding todo:", title)
	todo := &Todo{Title: title}
	t.Todos = append(t.Todos, *todo)
	return c.JSON(http.StatusCreated, todo)
}
