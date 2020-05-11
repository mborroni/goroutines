package main

import (
	"github.com/mborroni/goroutines/integration/internal/handler"
	"sync"
)

func main() {
	handler := handler.NewHandler()
	handler.AddTask("task one")
	
	length := handler.GetTasksLen()
	ch := make(chan string, length)
	var wg sync.WaitGroup
	
	handler.ExecutePendingTasks(ch, &wg)
}
