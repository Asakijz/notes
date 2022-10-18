package lesson05

import "testing"

//命名写错了，懒得改了
//本节是关于条件和循环的案例

func TestWhileLoop(t *testing.T) {
	//在Go中While循环写法如下所示
	n := 0
	for n < 5 {
		t.Log(n)
		n++
	}
}

func TestIfMultiSec(t *testing.T) {
	//可以先声明变量，用变量来判断结果
	if a := 1 == 1; a {
		t.Log("1==1")
	}
}

func TestSwitchMultiCase(t *testing.T) {
	//switch case写法
	//case后可跟多个选项
	for i := 0; i < 5; i++ {
		switch i {
		case 0, 2:
			t.Log(i, " is a Even number")
		case 1, 3:
			t.Log(i, " is a Odd number")
		default:
			t.Log(i, " is not a 0-3 number")
		}
	}
}

func TestSwitchCaseCondition(t *testing.T) {
	//在case中进行条件判断
	/*这种写法更像
		if
		elseif
		else
	的写法
	*/
	for i := 0; i < 5; i++ {
		switch {
		case i%2 == 0:
			t.Log(i, " is a Even number")
		case i%2 == 1:
			t.Log(i, " is a Odd number")
		default:
			t.Log(i, " is not a 0-3 number")
		}
	}
}
