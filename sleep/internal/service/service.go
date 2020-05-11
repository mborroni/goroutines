package service

import (
	"errors"
	"fmt"
)

//go:generate mockgen -source=service.go -destination=service_mock.go -package=service Service

var (
	ErrDuplicate = errors.New("task already exists")
)

type Service interface {
	GetTasks() []Task
	Add(task Task) error
	Run(task Task)
}

type service struct {
	tasks []Task
}

func NewService() *service {
	return &service{
		tasks: make([]Task, 0),
	}
}

func (s *service) GetTasks() []Task {
	return s.tasks
}

func (s *service) Add(task Task) error {
	if s.contains(task) {
		return ErrDuplicate
	}
	fmt.Println("[event:adding task][task:"+task+"]")
	s.tasks = append(s.tasks, task)
	return nil
}

func (s *service) Run(task Task) {
	fmt.Println("[event:executing task][task:"+task+"]")
}

func (s *service) contains(newTask Task) bool {
	for _, task := range s.tasks {
		if task == newTask {
			return true
		}
	}
	return false
}