package empty_interface

import (
	"fmt"
	"testing"
)

/*
	1. 空接口可以表示任何类型
	2. 通过断言来将空接口转换为执行类型
*/
func DoSomething(p interface{}) {
	switch v := p.(type) {
	case int:
		fmt.Println("Integer", v)
	case string:
		fmt.Println("string", v)
	default:
		fmt.Println("Unknown Type")
	}
}

func DoSomething2(p interface{}) {
	if i, ok := p.(int); ok {
		fmt.Println("Integer", i)
		return
	}
	if i, ok := p.(string); ok {
		fmt.Println("string", i)
		return
	}
	fmt.Println("Unknown Type")
}

func TestEmptyInterfaceAssertion(t *testing.T) {
	DoSomething("1")
	DoSomething(1)
	DoSomething(1.0)

	DoSomething2(1)
	DoSomething2("1")
	DoSomething2(1.0)
}
