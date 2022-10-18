package array_test

import "testing"

func TestArrayInit(t *testing.T) {
	var arr [3]int
	arr1 := [4]int{1, 2, 3, 4}
	arr2 := [...]int{1, 3, 4, 5} //可以用...来自动推导长度
	t.Log(arr[1], arr[2])
	t.Log(arr1)
	t.Log(arr2)
	arr1[1] = 5 //数组元素的访问通过下标，修改也是
	t.Log(arr1)
}

func TestArrayTravel(t *testing.T) {
	arr3 := [...]int{1, 3, 4, 5}
	//传统循环方法，通过len函数获取长度循环
	//for i := 0; i < len(arr3); i++ {
	//	t.Log(arr3[i])
	//}
	//如果不想接收一个值，可以用_来占位
	for index, value := range arr3 {
		t.Log(index, value)
	}
}

func TestArraySection(t *testing.T) {
	arr3 := [...]int{1, 2, 3, 4, 5}
	arr3_sec := arr3[:3] //可以通过下标来快速取值
	t.Log(arr3_sec)
}
