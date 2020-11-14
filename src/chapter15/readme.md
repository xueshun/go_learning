# 构建复用的包

## package
1. 基本复用模块单元
    以首字母大写来表示可被包外代码访问 (小写是不可以)
```go
func GetFibonacciSeries(n int) []int {
	ret := []int{1, 1}
	for i := 2; i < n; i++ {
		ret = append(ret, ret[i-2]+ret[i-1])
	}
	return ret
}
```

2. 代码的package可以和所在的目录不一致
3. 同一目录里的Go代码的package要保持一致

## 外部package
1. 通过go get 来获取远程依赖
    go get -u 强制从网络更新远程依赖
```go
 go  get -u  https://github.com/easierway/concurrent_map
```

## init 方法

1. 在 main 被强制执行前，所有依赖的 package 的 init 方法都会被执行
2. 不同包的 init 函数按照导入的依赖关系决定执行顺序
3. 每个包可以有多个 init 函数
4. 报的每个源文件也可以有多个 init 函数，这点比较特殊