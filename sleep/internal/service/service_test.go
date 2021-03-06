package service

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func NewMock() *service {
	return &service{
		tasks: make([]Task, 0),
	}
}

func TestService_Add(t *testing.T) {
	ass := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	
	service := NewMock()
	err := service.Add("print essay")
	ass.NoError(err)
}

func TestService_Add_ErrorOnDuplicatedTask(t *testing.T) {
	ass := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	
	service := NewMock()
	err := service.Add("print essay")
	ass.NoError(err)
	err = service.Add("print essay")
	ass.Error(err)
}

func TestService_GetTasks(t *testing.T) {
	ass := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	
	service := NewMock()
	tasks := service.GetTasks()
	ass.Len(tasks, 0)
}