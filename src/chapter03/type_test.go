package chapter03

import (
	"testing"
	_ "testing"
)

type MyInt int64

/*
Go 语言不允许隐式类型转换
别名和原有类型也不能进行隐式转换
*/
func TestImplicit(t *testing.T) {
	var a int32 = 1
	var b int64
	b = int64(a)
	var c MyInt
	c = MyInt(b)
	t.Log(a, b, c)
}

/**
 不支持指针运算
	aPtr = aPtr + 1
*/
func TestPoint(t *testing.T) {
	a := 1
	aPtr := &a
	//aPtr = aPtr + 1
	t.Log(a, aPtr)
	t.Logf("%T %T", a, aPtr)
}

/*
string 是值类型，其默认的初始值为空字符串，而不是nil
*/
func TestString(t *testing.T) {
	var s string
	t.Log("*" + s + "*")
	t.Log(len(s))
}
