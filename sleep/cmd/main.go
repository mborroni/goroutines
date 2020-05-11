package main

import (
	"github.com/mborroni/goroutines/sleep/internal/handler"
	"time"
)

func main() {
	handler := handler.NewHandler()
	handler.AddTask("task one")
	handler.AddTask("task two")
	handler.AddTask("task three")
	
	handler.ExecutePendingTasks()
	time.Sleep(100*time.Millisecond)
}