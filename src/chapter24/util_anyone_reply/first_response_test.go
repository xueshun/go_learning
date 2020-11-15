package util_anyone_reply

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func runTask(taskId int) string {
	time.Sleep(time.Millisecond * 10)
	return fmt.Sprintf("the result is from %d", taskId)
}

func FirstResponse() string {
	numOfRunner := 10
	// channel 和 buffer_channel的区别
	// channel 生产者必须要等待消费者消费
	// buffer_channel 生产者只需要将消费放到channel中即可，不需要等待
	ch := make(chan string)
	// ch := make(chan string, numOfRunner)
	for i := 0; i < numOfRunner; i++ {
		go func(i int) {
			ret := runTask(i)
			ch <- ret
		}(i)
	}
	// 阻塞等待从channel中获取值
	return <-ch
}

func TestFirstResponse(t *testing.T) {
	// 打印当前运行多少协程
	t.Log("Before:", runtime.NumGoroutine())
	t.Log(FirstResponse())
	time.Sleep(time.Second * 1)
	t.Log("After : ", runtime.NumGoroutine())
}
