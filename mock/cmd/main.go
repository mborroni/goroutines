package main

import (
	"github.com/mborroni/goroutines/mock/internal/handler"
)

func main() {
	handler := handler.NewHandler()
	handler.AddTask("task one")
	handler.AddTask("task two")
	handler.AddTask("task three")
	
	handler.ExecutePendingTasks()
}
