package slice_test

import "testing"

func TestSliceInit(t *testing.T) {
	var s0 []int //切片和数组的区别在于--->切片是可变长度的
	t.Log(len(s0), cap(s0))
	s0 = append(s0, 1)
	t.Log(len(s0), cap(s0))

	//切片的声明
	s1 := []int{1, 3, 4, 5} //短变量声明
	t.Log(len(s1), cap(s1))

	s2 := make([]int, 3, 5) //make声明，第一个参数为类型，第二个参数为len长度，第三个参数为cap容量大小
	//s2声明了一个int类型的切片，长度为3，容量为5
	t.Log(len(s2), cap(s2))
	t.Log(s2[0], s2[1], s2[2]) //len为我们可访问的元素
	//t.Log(s2[3])               //cap为切片的容量，超过len会下标越界无法访问
	s2 = append(s2, 1)
	t.Log(s2[0], s2[1], s2[2], s2[3])
	t.Log(len(s2), cap(s2))
}

func TestSliceGrowing(t *testing.T) {
	s := []int{}
	// 默认的 len 和 cap 均为空
	// 后续 for 循环添加数值时，cap 值会在原有基础上*2
	for i := 0; i < 10; i++ {
		s = append(s, i)
		t.Log(len(s), cap(s))
	}
}

func TestSliceShareMemory(t *testing.T) {
	// 截取切片后，len长度为所截取元素的长度，cap容量为所截取的起始位置至原始切片的最后一位
	month := []string{"Jan", "Feb", "Mon", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}
	Q2 := month[3:6]
	t.Log(Q2, len(Q2), cap(Q2))
	summer := month[5:8]
	t.Log(summer, len(summer), cap(summer))
	summer[0] = "Unknown"
	t.Log(Q2) //因为是共享一片连续存储空间，所以对summer的修改也会影响其他切片的值
	t.Log(month)
}

func TestSliceComparing(t *testing.T) {
	a := []int{1, 2, 3, 4}
	b := []int{1, 2, 3, 4}
	//两个切片只能和nil进行比较
	//if a == b {
	//	t.Log("Equal")
	//}
	if a != nil {
		t.Log("a is a Slice")
	}
	if b != nil {
		t.Log("b is a Slice")
	}
}
