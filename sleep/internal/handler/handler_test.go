package handler

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/mborroni/goroutines/sleep/internal/service"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
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

func TestHandler_ExecutePendingTasks(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	handler := NewMock(ctrl)
	handler.service.(*service.MockService).EXPECT().GetTasks().Return([]service.Task{"task"})
	handler.service.(*service.MockService).EXPECT().Run(gomock.Any())

	handler.ExecutePendingTasks()
	time.Sleep(105*time.Millisecond)
}
