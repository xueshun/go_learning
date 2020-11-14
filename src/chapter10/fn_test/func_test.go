package fn_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

/*
	返回多个Value
*/
func returnMultiValues() (int, int) {
	return rand.Intn(10), rand.Intn(20)
}

/*
	监控一个函数的执行时间
*/
func timeSpent(inner func(op int) int) func(op int) int {
	return func(op int) int {
		start := time.Now()
		ret := inner(op)
		fmt.Println("time spent:", time.Since(start).Seconds())
		return ret
	}
}

/*
	执行耗时函数
*/
func slowFn(op int) int {
	time.Sleep(time.Second * 1)
	return op
}

func TestFn(t *testing.T) {
	a, _ := returnMultiValues()
	t.Logf("returnMultiValues: %d", a)
	tsFn := timeSpent(slowFn)
	t.Log(tsFn(10))
}

/*
	可变参数
*/
func Sum(ops ...int) int {
	ret := 0
	for _, op := range ops {
		ret += op
	}
	return ret
}

func TestVarParam(t *testing.T) {
	t.Log(Sum(1, 2, 3, 4))
	t.Log(Sum(1, 2, 3, 4, 5))
}

/*
	延时运行
*/
func Clear() {
	fmt.Println("Clear resources.")
}

func TestDefer(t *testing.T) {
	defer Clear() // 释放资源，锁
	fmt.Println("Start ")
	// 异常中断
	panic("err")
}
