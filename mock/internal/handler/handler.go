package handler

import (
	"fmt"
	"github.com/mborroni/goroutines/mock/internal/service"
	"sync"
)

//go:generate mockgen -source=handler.go -destination=handler_mock.go -package=handler Handler

type Handler interface {
	AddTask(task service.Task) error
	ExecutePendingTasks()
}

type handler struct {
	service service.Service
}

func NewHandler() *handler {
	return &handler{
		service: service.NewService(),
	}
}

func (h *handler) AddTask(task service.Task) error {
	return h.service.Add(task)
}

func (h *handler) ExecutePendingTasks() {
	var wg sync.WaitGroup
	ch := make(chan string, len(h.service.GetTasks()))
	for _, task := range h.service.GetTasks() {
		wg.Add(1)
		go h.service.Run(task, ch, &wg)
		wg.Wait()
	}
	close(ch)
	
	for result := range ch {
		fmt.Println(result)
	}
}

