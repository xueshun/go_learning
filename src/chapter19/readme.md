
# 多路选择和超时控制

## Select 

```go
select {
case ret := <- retCh1:
    t.Logf("result %s", ret)
case ret := <- retCh2:
    t.Logf("result %s", ret)
default:
    t.Error("No ont returned")
}
```

超时控制
```go
select {
case ret:= <- retCh:
    t.Logf("result %s", ret)
case <- time.After(time.Second *1)
    t.Error("time out")
}
```