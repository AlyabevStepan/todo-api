package main

import (
	"log"
	"github.com/AlyabevStepan/todo-api"
	"github.com/AlyabevStepan/todo-api/pkg/handler"
)

func main() {
	handlers := new(handler.Handler)
	srv := new(todo.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("%s", err.Error())
	}
}
