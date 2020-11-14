
# GO 编写结构话程序

## 循环

```go
for i := 0; i < 5; i++
```

while 条件循环
```go
n := 0
for n < 5 {
    n++
    fmt.Println(n)
}
```

无限循环
```go
n := 0
for {
}
```

## if 条件

```go

```


## Switch

与其他语言的差异
1. 条件表达式不限制为常量或整数
2. 单个case中，可以出现多个结果选线，使用时使用逗号分隔
3. case 不需要break
4. 可以不设定switch之后的条件表达式，在此情况下，整个switch结构与多个if...else...的逻辑作用等同
