package main

import "fmt"

//函数功能：求和
// 函数的名字：getSum 参数为空
// getSum函数的返回值为一个函数，这个函数的参数是一个int类型的参数，返回值也是int类型
func getSum() func(int) int {
	var sum int = 0
	return func(num int) int {
		sum += num
		return sum
	}
}

// 闭包：返回的匿名函数+匿名函数外的变量num
// 匿名函数中引用的那个变量会一直保存在内存中，可以一直使用

func main() {
	f := getSum()
	fmt.Println(f(1))
	fmt.Println(f(2))
}

/*
闭包的本质：
闭包本质依旧是一个匿名函数，只是这个函数引入外界的变量/参数
匿名函数+引用的变量/参数 = 闭包

闭包特点：
1、返回的是一个匿名函数，但是这个匿名函数引用到函数外的变量/参数，因此这个匿名函数就和外界变量/参数形成一个整体，构成闭包
2、闭包中使用的外界参数/变量会一直保存在内存中，所以一直可以使用（意味着闭包不可滥用，对内存消耗很大）
*/
