package _select

import (
	taskService "chapter18/service"
	"testing"
	"time"
)

/*
	测试多渠道选择
*/
func TestSelect(t *testing.T) {
	select {
	case ret := <-taskService.AsyncService():
		t.Logf(ret)
	case <-time.After(time.Millisecond * 40):
		t.Error("time out")
	}
}
