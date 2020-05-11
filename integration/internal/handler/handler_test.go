package handler

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/mborroni/goroutines/integration/internal/service"
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
)

func NewMock(ctrl *gomock.Controller) *handler {
	return &handler{
		service: service.NewMockService(ctrl),
	}
}

func TestHandler_AddTask(t *testing.T) {
	ass := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	
	handler := NewMock(ctrl)
	handler.service.(*service.MockService).EXPECT().Add(gomock.Any()).Return(nil)
	err := handler.AddTask("do stuff")
	ass.NoError(err)
}

func TestHandler_AddTask_Error(t *testing.T) {
	ass := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	
	handler := NewMock(ctrl)
	handler.service.(*service.MockService).EXPECT().Add(gomock.Any()).Return(errors.New("error"))
	err := handler.AddTask("do stuff")
	ass.Error(err)
}

func TestHandler_GetTasksLen(t *testing.T) {
	ass := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	
	handler := NewMock(ctrl)
	handler.service.(*service.MockService).EXPECT().GetTasks().Return(nil)
	len := handler.GetTasksLen()
	ass.Equal(0, len)
}

func TestHandler_ExecutePendingTasks(t *testing.T) {
	handler := NewHandler()
	handler.AddTask("write something")
	handler.AddTask("test something")
	handler.AddTask("deploy something")
	
	ch := make(chan string, 3)
	var wg sync.WaitGroup
	
	handler.ExecutePendingTasks(ch, &wg)
}
