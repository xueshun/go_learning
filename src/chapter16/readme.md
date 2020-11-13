# 协程机制

## Thread vs Coroutine

1. 创建时默认的stack的大小
    - JDK以后的 Java Thread Stack 默认为 1M
    - Coroutine 的 Stack 默认大小为 2k
2. 和KSE(Kernel Space Entity)的对应关系
    - Java Thread 1:1
    - Coroutine M:N

![图 1](../../images/82fa198e5c09c107653ee32f19b70483db73c664ac8b78edeb33dee76896a3f0.png)  
![图 2](../../images/c932bc80cafa8a5e7d4d2ebed2339361a4c07f57986b24ad22aa9776ed9560a5.png)  
