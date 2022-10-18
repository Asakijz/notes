package lesson10

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func ReturnMultiValue() (int, int) {
	return rand.Intn(10), rand.Intn(20)
}

func timeSpent(inner func(a int) int) func(op int) int {
	return func(n int) int { // 2
		start := time.Now()                                     // 4
		ret := inner(n)                                         // 5
		fmt.Println("time spent:", time.Since(start).Seconds()) // 8
		return ret                                              // 9
	}
}

func slowFun(op int) int {
	time.Sleep(time.Second * 1) // 6
	return op                   // 7
}

func TestFn(t *testing.T) {
	// firstFunc
	a, _ := ReturnMultiValue()
	t.Logf("first random number is:%d", a)
	// secondFunc
	// 计算函数内所耗时长
	tsSF := timeSpent(slowFun) // 1
	t.Log(tsSF(10))            // 3
}

// 可变参数写法 ...后面跟类型
func Sum(ops ...int) int {
	ret := 0
	for _, v := range ops {
		ret += v
	}
	return ret
}

func TestVarParam(t *testing.T) {
	t.Log(Sum(1, 2, 3, 4))
	t.Log(Sum(1, 2, 3, 4, 5))
}

// 延迟调用
func Clean() {
	fmt.Println("clean resources.")
}

func TestDefer(t *testing.T) {
	defer Clean()
	t.Log("Start")
	panic("err")       // 在 panic 后 defer 的语句也能执行。通常用来关闭锁或者释放资源
	fmt.Println("End") // 在 panic 后的语句是无法执行下去的
}
