package handler

import (
	"fmt"
	"github.com/mborroni/goroutines/integration/internal/service"
	"sync"
)

//go:generate mockgen -source=handler.go -destination=handler_mock.go -package=handler Handler

type Handler interface {
	AddTask(task service.Task) error
	GetTasksLen() int
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

func (h *handler) GetTasksLen() int {
	return len(h.service.GetTasks())
}

func (h *handler) ExecutePendingTasks(ch chan string, wg *sync.WaitGroup) {
	for _, task := range h.service.GetTasks() {
		wg.Add(1)
		go h.service.Run(task, ch, wg)
		wg.Wait()
	}
	close(ch)
	
	for result := range ch {
		fmt.Println(result)
	}
}

