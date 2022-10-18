package lesson19

import (
	"fmt"
	"testing"
	"time"
)

func service() string {
	time.Sleep(time.Millisecond * 500)
	return "Done"
}

func AsyncService() chan string {
	//retCh := make(chan string)
	retCh := make(chan string, 1) //第二个参数声明 cap 容量，这样不会造成阻塞
	go func() {
		ret := service()
		fmt.Println("returned result")
		retCh <- ret
		fmt.Println("service exited")
	}()
	return retCh
}

func TestAsyncService(t *testing.T) {
	// select 和 switch 不同，并不是按照指定的 case 顺序执行，而是哪个协程满足条件就执行哪个
	select {
	case ret := <-AsyncService():
		t.Log(ret)
	case <-time.After(time.Millisecond * 100):
		t.Error("time out")
	}

}
