// 在Go中，函数也是一种数据类型，可以赋值给变量，则该函数就是一个函数类型的变量了。可以作为函数的参数和返回值。
package main
import "fmt"

// 定义函数
func add(a, b int) int {
	return a + b
}

//定义一个函数，把另一个函数作为形参
func calc(a, b int, op func(int, int) int) int { //  op func(int, int) int 是一个函数类型的参数 
	return op(a, b)
}

a := add
fmt.Printf("a的类型：%T\n", a) // a的类型：func(int, int) int
fmt.Println(a(1, 2)) // 3

func main() {
	// 调用函数
	ret := calc(10, 20, add)
	fmt.Println(ret) // 30
}

