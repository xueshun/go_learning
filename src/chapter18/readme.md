# CSP (Communicating sequential process)

## Actor

![图 1](../../images/957a887be35d28b4f485be2330c70f2ae2b776e1305504ec81f1a9a3253cee51.png)  

## CSP vs Actor

1. 和Actor的直接通讯不同，CSP模式则是通过Channel进行通讯的，耦合更松一些

2. Go中channel是有容量限制并且独立处理Groutine,  而如Erlang, Actor中的mailbox容量是无限的，接收进程也总是被动处理消息

![图 2](../../images/682a928fc7b80c7a546d7a401589f539afd774cd8b00bb3d1c206ed4c948657e.png)  

## channel

![图 3](../../images/a82920d14e8840b9a460870c203e4f2513ca46dec0d268af5b4aeb6e6cf78f7d.png)  

