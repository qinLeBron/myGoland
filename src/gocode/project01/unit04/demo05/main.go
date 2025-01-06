package main

import "fmt"

func main() {
	fmt.Println(add(10, 20))
}

func add(num1 int, num2 int) int {
	// 在Golang中，程序遇到defer关键字，不会立即执行defer后的语句
	// 而是将defer后的语句压入一个栈中，然后继续执行函数后面的语句
	defer fmt.Println("num1 = ", num1)
	defer fmt.Println("num2 = ", num2)
	num1 += 50
	num2 += 60

	// 栈的特点：先进后出
	// 在函数执行完毕后，从栈中取出语句开始执行，按照先进后出的规则执行语句
	// 遇到defer关键字，会将后面的代码语句压入栈中，也会将相关的值同时拷贝入栈中，不糊随着函数后面的变化而变化
	var sum int = num1 + num2
	fmt.Println("sum = ", sum)
	return sum
}

/*
defer应用场景：
比如你想关闭某个使用的资源，在使用的时候直接随手defer，因为defer有延迟执行机制（函数执行完再执行defer压入栈的语句），
所以用完随手写了关闭，比较省心省事
*/
