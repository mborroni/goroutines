package handler

import (
	"github.com/mborroni/goroutines/sleep/internal/service"
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
	for _, task := range h.service.GetTasks() {
		go h.service.Run(task)
	}
}

