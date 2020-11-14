package customer_type

import (
	"fmt"
	"testing"
	"time"
)

func timeSpent(inner func(op int) int) func(op int) int {
	return func(op int) int {
		start := time.Now()
		ret := inner(op)
		fmt.Print("time Spent:", time.Since(start).Seconds())
		return ret
	}
}

type IntConv func(op int) int

func timeSpentCustomer(inner IntConv) IntConv {
	return func(op int) int {
		start := time.Now()
		ret := inner(op)
		fmt.Println("time spent:", time.Since(start).Seconds())
		return ret
	}
}

func slowFn(op int) int {
	time.Sleep(time.Second * 1)
	return op
}

func TestFn(t *testing.T) {
	tsFn := timeSpentCustomer(slowFn)
	t.Log(tsFn(10))
}
