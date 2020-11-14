package cancel_by_close

import (
	"fmt"
	"testing"
	"time"
)

/*
	获取取消通知
*/
func isCanceled(cancelChannel chan struct{}) bool {
	select {
	case <-cancelChannel:
		return true
	default:
		return false
	}
}

/*
	发送取消消息
*/
func cancel_1(cancelChannel chan struct{}) {
	cancelChannel <- struct{}{}
}

/*
	发送取消消息
*/
func cannel_2(cancelChannel chan struct{}) {
	close(cancelChannel)
}

func TestCannel(t *testing.T) {
	cancelChannel := make(chan struct{}, 0)
	for i := 0; i < 5; i++ {
		go func(i int, cancelCh chan struct{}) {
			for {
				if isCanceled(cancelCh) {
					break
				}
				time.Sleep(time.Millisecond * 5)
			}
			fmt.Println(i, "Canceled")
		}(i, cancelChannel)
	}
	// 发送一个channel关闭信号，又是我们是不知道有几个channel需要关闭，
	// 而且对每个channel都发送关闭信号，那么代码的耦合度将会提高
	cancel_1(cancelChannel)
	time.Sleep(time.Second * 1)
}
