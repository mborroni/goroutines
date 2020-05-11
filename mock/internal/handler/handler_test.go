package handler

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/mborroni/goroutines/mock/internal/service"
	"github.com/stretchr/testify/assert"
	"sync"
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
	
	tasks := getTasksList()
	
	handler := NewMock(ctrl)
	handler.service.(*service.MockService).EXPECT().GetTasks().AnyTimes().Return(tasks)
	handler.service.(*service.MockService).EXPECT().Run(gomock.Any(), gomock.Any(), gomock.Any()).Do(func(task service.Task, ch chan string, wg *sync.WaitGroup) {
		ch <- "[event:executing task][task:"+string(task)+"]"
		wg.Done()
		
		done := make(chan struct{})
		go func() {
			wg.Wait()
			close(done)
			/*
			    value, ok := <- done
				indicates whether the read off the channel was a value generated by a write elsewhere in the process,
			    or a default value generated from a closed channel
			*/
		}()
		
		select {
		case <-done:
			/* time out avoid blocking */
		case <-time.After(10*time.Second):
		}
	}).AnyTimes()
	
	handler.ExecutePendingTasks()
}

func getTasksList() []service.Task {
	return []service.Task{
		"write something",
		"test something",
		"deploy something",
	}
}
