package until_all_done

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func runTask(taskId int) string {
	time.Sleep(time.Millisecond * 10)
	return fmt.Sprintf("the result is from%d", taskId)
}

func AllResponse() string {
	numOfRunner := 10
	ch := make(chan string, numOfRunner)
	for i := 0; i < numOfRunner; i++ {
		go func(i int) {
			ret := runTask(i)
			ch <- ret
		}(i)
	}
	finalRet := ""
	for j := 0; j < numOfRunner; j++ {
		finalRet += <-ch + "\n"
	}
	return finalRet
}

func AllResponseByWg() string {
	// todo by waitGroup
	return ""
}

func TestAllResponse(t *testing.T) {
	t.Log("Before: ", runtime.NumGoroutine())
	t.Log(AllResponse())
	time.Sleep(time.Second * 1)
	t.Log("After:", runtime.NumGoroutine())
}
