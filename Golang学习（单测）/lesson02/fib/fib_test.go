package fib

import (
	"fmt"
	"testing"
)

//斐波那契数列（递归函数）
func TestFibList(t *testing.T) {
	a := 1
	b := 1
	fmt.Print(a)
	for i := 0; i < 5; i++ {
		fmt.Print(" ", b)
		tmp := a
		a = b
		b = tmp + a
	}
	fmt.Println()
}
