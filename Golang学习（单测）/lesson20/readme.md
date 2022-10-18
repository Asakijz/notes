## channel 的关闭

* 向关闭的 channel 发送数据，会导致 panic
* v,ok <- ch;ok 为 bool 值，true 表示正常接收，false 表示通道关闭
* 所有的 channel 接收者都会在 channel 关闭时，立刻从阻塞等待中返回且上述 ok 值为 false。这个广播机制常被利用，进行向多个订阅者同时发送信号。如：退出信号

## Context

* 根 Context：通过 context.Background() 创建
* 子 Context：context.WithCancel(parentContext) 创建
  * ctx,cancel := context.WithCancel(context.Background())
* 当前 Context 被取消时，基于他的子 context 都会被取消
* 接收取消通知 <- ctx.Done()