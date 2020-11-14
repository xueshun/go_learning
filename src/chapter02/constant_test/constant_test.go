package constant_test

import "testing"

// 累加
const (
	Monday = 1 + iota
	Tuesday
	Wednesday
)

// 位运算
const (
	Readable = 1 << iota
	Writeable
	Executable
)

func TestConstantTray(t *testing.T) {
	t.Log(Monday, Tuesday)
}

func TestConstantTry1(t *testing.T) {
	a := 1 // 0001
	t.Log(a&Readable == Readable, a&Writeable == Writeable, a&Executable == Executable)
}
