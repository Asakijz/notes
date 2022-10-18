package lesson14

import (
	"errors"
	"fmt"
	"testing"
)

func TestPanicVsExit(t *testing.T) {
	//defer func() {
	//	fmt.Println("Finally!")
	//}()
	/*个人建议最好不要这么写，很容易导致僵尸进程（即资源消耗殆尽）
	最好的方法是直接让它崩溃*/
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recovered from", err)
		}
	}()
	fmt.Println("Start")
	panic(errors.New("Something wrong!"))
	//os.Exit(-1) // 如果用的是 os.Exit，那么 defer 语句并不会执行
}
