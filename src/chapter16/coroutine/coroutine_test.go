package coroutine

import (
	"fmt"
	"testing"
	"time"
)

func TestCoroutine(t *testing.T) {
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println(i)
		}(i)
	}
	time.Sleep(time.Millisecond * 50)
}

func TestCoroutine1(t *testing.T) {
	for i := 0; i < 10; i++ {
		// 这些写法是错误的，i为公用的地址
		go func() {
			fmt.Println(i)
		}()
	}
	time.Sleep(time.Millisecond * 50)
}
