# Context 与任务取消

## 关联任务取消

TODO

![图 4](../../images/f9f1fcc2c21dfb362b7e418c8305f82877857105dab35d890b5a0c7cf39a4f0a.png)  


## Context

1. 根Context：通过context.Background() 创建
2. 子Context：context.WithCancel(parentContext) 创建
    - ctx, cancel := context.WithCancel(context.Background)
3. 当前Context被取消时，基于他的子context都会被取消
4， 接收取消通知 <- ctx.Done()