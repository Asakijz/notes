package main

import (
	"fmt"
	"os"
)

func main() {
	//os.Args捕获命令行参数
	//fmt.Println(os.Args)

	//如果命令行参数长度大于1，打印hello world，并跟上命令行参数
	if len(os.Args) > 1 {
		fmt.Println("hello world", os.Args[1])
	}

	//打印hello world
	//fmt.Println("hello world")

	//捕获异常退出值
	//os.Exit(-1)
}
