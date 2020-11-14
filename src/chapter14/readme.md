
# Go错误机制 Panic和Recover

## Go错误机制

1. 没有错误机制
2. error类型实现了error接口
```go
type error interface{
    Error() string
}
```
3. 可以通过errors.New 来快速创建错误实例
```go
errors.New("n must be in the range[0,100]")
```

### 实践

**尽早失败，避免嵌套**

定义不同的错误类型，以便判断错误类型
```go
var LessThanTwoError error = errors.New("n must be greater than 2")
var GreateThanHundredError error = errors.New("n must be less than 100")

func TestFibonacci(t *testing.T){
   if n < 2 {
   		return nil, LessThanTwoError
   	}
   	if n > 10 0 {
   		return nil, LargerThenHundredError
   	}
}
```

## panic 和 recover

### panic
1. panic 用于不可以恢复的错误
2. panic 退出前会执行defer指定的内容

### panic vs os.Exit

1. os.Exit 退出时不会调用defer指定的函数
2. os.Exit 退出时不输出当前调用栈信息

### recover

错误恢复
```go
if err:= recover(); err != nil{
	fmt.Println("recovered from", err)
}
```

