package series

import "fmt"

// init 函数可以有多个
func init() {
	fmt.Println("first init")
}

func init() {
	fmt.Println("second init")
}

// 要将包名首字母大写对外暴露，不然包外无法访问调用
/*func square(n int) int {
	return n * n
}*/

func Square(n int) int {
	return n * n
}

func GetFibonacciSerie(n int) []int {
	ret := []int{1, 1}
	for i := 2; i < n; i++ {
		ret = append(ret, ret[i-2]+ret[i-1])
	}
	return ret
}
