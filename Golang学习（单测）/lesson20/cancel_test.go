package lesson20

import (
	"fmt"
	"testing"
	"time"
)

// 判断是否有取消任务的信号
func isCancelled(cancelChan chan struct{}) bool {
	select {
	case <-cancelChan: // 当 channel 接收消息时返回 true，TestCancel 函数直接 break，打印当前循环的 i 值
		return true
	default:
		return false
	}
}

//发送信号的个数需要与关闭的协程数量一致
func cancel_1(cancelChan chan struct{}) {
	cancelChan <- struct{}{}
}

// 直接关闭所有协程
func cancel_2(cancelChan chan struct{}) {
	// channel 被 close 后，channel 仍然可读，不但可以读取已发出的数据，还可以不断读取零值
	// 但是如果这个 channel 是通过 range 来读取的，那么在 channel 被关闭后，for 循环会退出
	close(cancelChan)
}

func TestCancel(t *testing.T) {
	cancelChan := make(chan struct{}, 0)
	for i := 0; i < 5; i++ {
		go func(i int, cancelCh chan struct{}) {
			for {
				if isCancelled(cancelCh) {
					break
				}
				time.Sleep(time.Millisecond * 5)
			}
			fmt.Println(i, "Cancel")
		}(i, cancelChan)
	}
	cancel_2(cancelChan)
	time.Sleep(time.Second)
}
