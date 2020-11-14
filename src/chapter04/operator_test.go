package chapter04

import "testing"

const (
	Readable = 1 << iota
	Writable
	Executable
)

/**
比较数组注意点：
	1. 相同维数且含有相同个数元素的数组才可以比较
	2. 每个元素都相同才相等
*/
func TestCompareArray(t *testing.T) {
	a := [...]int{1, 2, 3, 4}
	b := [...]int{1, 2, 3, 4}
	c := [...]int{1, 2, 3, 5}
	//d := [...]int{1, 2, 3, 4, 5}
	t.Log(a == b)
	t.Log(a == c)
	//t.Log(a == d)
}

func TestBitClear(t *testing.T) {
	a := 7            // 0111
	a = a &^ Readable // 0111 &^ 0001  => 0110
	t.Log(a)
	a = a & ^Executable // 0110 &^ 0111 =>
	t.Log(a)
	t.Log(a&Readable == Readable,
		a&Writable == Writable,
		a&Executable == Executable)
}
