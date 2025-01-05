package main

import (
	"fmt"
	"gocode/project01/unit04/demo02/testutils" //导包时会执行testutils的变量和init函数
)

var num int = test()

func test() int {
	fmt.Println("test函数被执行了。。。")
	return 100
}

/*
init函数，初始化函数，可以用来进行一些初始化操作，
每一个源文件都可以包含一个init函数，该函数会在main函数执行之前被Go运行框架调用
*/
func init() {
	fmt.Println("init函数被执行了。。。")
}

func main() {
	fmt.Println("main函数被执行了。。。")
	fmt.Println("Name = ", testutils.Name)
}
