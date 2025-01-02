// Go支持自定义数据类型   基本语法如下：
// type 自定义数据类型名 数据类型
package main
import "fmt"

// 定义一个函数类型
type myFuncType func(int, int) int // myFuncType 是一个函数类型

// 定义函数
func add(a, b int) int {
	return a + b
}

//定义一个函数，把另一个函数作为形参
func calc(a, b int, op myFuncType) int { //  op myFuncType 是一个函数类型的参数 
	return op(a, b)
}

func main() {
	// 调用函数
	ret := calc(10, 20, add)
	fmt.Println(ret) // 30
}
