# channel 的关闭和广播

## channel 的关闭

1. 向关闭的channel发送数据，会导致panic
2. v,ok := <- ch, ok 为bool值，true表示正常接收，false表示通道关闭
3. 所有的channel接收者都会在channel关闭时，立刻从阻塞等待中返回且上述ok值为false。这个
广播机制常被利用，进行多个订阅者同时发送信息。如：退出信号