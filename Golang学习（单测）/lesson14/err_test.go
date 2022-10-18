package lesson14

import (
	"errors"
	"fmt"
	"testing"
)

//定义不同的错误变量，易于判断错误的类型
var lessThan = errors.New("n is less than 2")
var lagerThan = errors.New("n is large than 100")

func GetFibonacci(n int) ([]int, error) {
	//if n < 2 || n > 100 {
	//	return nil, errors.New("n should be in [0,100]")
	//}
	if n < 2 {
		return nil, lessThan
	}
	if n > 100 {
		return nil, lagerThan
	}
	fibList := []int{1, 1}
	for i := 2; i < n; i++ {
		fibList = append(fibList, fibList[i-2]+fibList[i-1])
	}
	return fibList, nil
}

func TestGetFibonacci(t *testing.T) {
	if v, err := GetFibonacci(110); err != nil {
		if err == lessThan {
			fmt.Println("it is less")
		}
		if err == lagerThan {
			fmt.Println("it is large")
		}
		t.Error(err)
	} else {
		t.Log(v)
	}
}
