package service

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
	"time"
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

func TestService_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	
	ch := make(chan string, 1)
	var wg sync.WaitGroup
	wg.Add(1)
	
	service := NewMockService(ctrl)
	service.EXPECT().Run(gomock.Any(), gomock.Any(), gomock.Any()).Do(func(task Task, ch chan string, wg *sync.WaitGroup) {
		defer wg.Done()
		ch <- "print essay"
		
		done := make(chan struct{})
		go func() {
			wg.Wait()
			close(done)
		}()
		select {
		case <-done:
		case <-time.After(5*time.Second):
		}
	})
	service.Run("print essay", ch, &wg)
	wg.Wait()
}