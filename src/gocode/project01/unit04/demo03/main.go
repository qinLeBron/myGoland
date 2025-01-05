package main

import "fmt"

var Func01 = func(num1 int, num2 int) int {
	return num1 * num2
}

func main() {
	//定义匿名函数：定义的同时调用
	result := func(num1 int, num2 int) int {
		return num1 + num2
	}(10, 20)

	fmt.Println(result)

	// 将匿名函数赋值给一个变量，这个变量实际就是函数类型的变量
	// sub等价于匿名函数
	sub := func(num1 int, num2 int) int {
		return num1 - num2
	}

	// 直接调用sub就是调用匿名函数了
	result02 := sub(20, 70)
	fmt.Println(result02)

	result03 := Func01(3, 4)
	fmt.Println(result03)
}
