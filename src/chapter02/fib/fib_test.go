package fib

import "testing"

/*
	1. 测试文件XXX_test.go , 必须以test结尾
	2. 测试方法TestXXX , 测试方法必须以Test开头
*/
func TestFibTest(t *testing.T) {
	var a int = 1
	var b int = 1
	t.Log(a)
	for i := 0; i < 5; i++ {
		t.Log(" ", b)
		tmp := a
		a = b
		b = tmp + a
	}
	t.Log(a, b)
}

func TestFib1Test(t *testing.T) {
	// 常量赋值
	var (
		a = 1
		b = 1
	)
	t.Log(a, b)
}

func TestFib2Test(t *testing.T) {
	// 常量赋值
	a := 1
	b := 1
	t.Log(a, b)
}

/**
测试两值交换
*/
func TestExchange(t *testing.T) {
	a := 1
	b := 2
	tmp := a
	a = b
	b = tmp
	t.Log(a, b)
}

func Test1Exchange(t *testing.T) {
	a := 1
	b := 2
	a, b = b, a
	t.Log(a, b)
}
