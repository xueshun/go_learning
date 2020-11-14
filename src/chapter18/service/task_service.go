package service

import (
	"fmt"
	"time"
)

func Service() string {
	// 模拟业务操作
	time.Sleep(time.Millisecond * 50)
	return "Service Done"
}

func OtherService() {
	fmt.Println("Working on Something else")
	time.Sleep(time.Millisecond * 100)
	fmt.Println("Task Is Done")
}

func AsyncService() chan string {
	// channel buffer_channel
	retCh := make(chan string, 1)
	go func() {
		ret := Service()
		fmt.Println("returned result.")
		retCh <- ret
		fmt.Println("service exited")
	}()
	return retCh
}
