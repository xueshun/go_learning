package string

import "testing"

func TestString(t *testing.T) {
	var s string
	t.Log(s) // 初始化位默认零值""
	s = "hello"
	t.Log(len(s))
	// s[1] = '3' string是不可变的byte slice
	s = "\xE4\xB8\xA5" // 可以存储任何二进制数据
	t.Log(s)
	s = "中"
	t.Log(len(s)) // len(s) byte数

	c := []rune(s)
	t.Log(len(c))
	t.Logf("中 unicode %x", c[0])
	t.Logf("中 UTF8 %x", s)
}

func TestStringToRune(t *testing.T) {
	s := "中华人民共和国"
	for _, c := range s {
		t.Logf("%[1]c, %[1]x", c)
	}
}
