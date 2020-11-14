package csp

import (
	taskService "chapter18/service"
	"fmt"
	"testing"
	"time"
)

func TestService(t *testing.T) {
	fmt.Println(taskService.Service())
	taskService.OtherService()
}

func TestAsyncService(t *testing.T) {
	retCh := taskService.AsyncService()
	taskService.OtherService()
	fmt.Println(<-retCh)
	time.Sleep(time.Second * 1)
}
