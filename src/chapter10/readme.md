
# 函数

## 函数一等公民

1. 可以有多个返回值
2. 所有参数都是值传递：slice,map,channel 会有引用传递的错觉
3. 函数可以作为变量的值
4. 函数可以作为参数和返回值

《计算机程序的构造和解释》

## 可变参数及defer

1. 可变参数

```go
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
```

2. defer 延迟执行，释放资源或锁

```go
func Clear(){
	fmt.Println("Clear resources.")
}

func TestDefer(t *testing.T){
	defer Clear() // 释放资源，锁
	fmt.Println("Start ")
	// 异常中断
	panic("err")
}
```
