package lesson22

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func runtask(id int) string {
	time.Sleep(time.Millisecond * 10)
	return fmt.Sprintf("the result is from %d", id)
}

func FirstResponse() string {
	numOfRunner := 10
	//ch := make(chan string) // 不推荐不声明容量，这样会导致协程泄漏
	ch := make(chan string, numOfRunner)
	for i := 0; i < numOfRunner; i++ {
		go func(i int) {
			ret := runtask(i)
			ch <- ret
		}(i)
	}
	return <-ch // 一但 ch 有值，会立马 return
}

func TestFirstResponse(t *testing.T) {
	t.Log("Before:", runtime.NumGoroutine())
	t.Log(FirstResponse())
	time.Sleep(time.Second)
	t.Log("After:", runtime.NumGoroutine())
}
