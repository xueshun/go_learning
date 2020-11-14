package channel_close

import (
	"fmt"
	"sync"
	"testing"
)

func dataProducer(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)
		// 如果channel 已经关闭，继续往channel存放，则会抛出异常
		// panic： send on closed channel
		// ch <- 11
		wg.Done()
	}()
}

func dataReceiver(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for {
			if data, ok := <-ch; ok {
				fmt.Println(data)
			} else {
				break
			}
		}
		// 如果channel中消费者消费数据多于生产者，则会返回默认值，如：0
		/*for i := 0; i < 11 ; i++ {
			data := <- ch
			fmt.Println(data)
		}*/
		wg.Done()
	}()
}

func TestCloseChannel(t *testing.T) {
	var wg sync.WaitGroup
	ch := make(chan int)
	wg.Add(1)
	dataProducer(ch, &wg)
	wg.Add(1)
	dataReceiver(ch, &wg)

	// 添加多个消费者
	// wg.Add(1)
	// dataReceiver(ch, &wg)
	wg.Wait()
}
