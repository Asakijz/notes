package lesson16

import (
	"fmt"
	"testing"
	"time"
)

func TestGoroutine(t *testing.T) {
	for i := 0; i < 10; i++ {
		// 因为都是值传递，这样写是把 i 复制了一份，每个 i 的地址不一样
		go func(i int) {
			fmt.Println(i)
		}(i)
		// 不推荐下面的写法，因为变量在这个协程中被共享了，共享变量存在竞争条件。所以最后会输出相同数字
		//go func() {
		//	fmt.Println(i)
		//}()
	}
	time.Sleep(time.Microsecond * 50)
}
