// 支持对函数返回值命名
package main
import "fmt"

// 定义函数
func test(a, b int) (sum int， sub int) { //  sum int， sub int 是对返回值命名
	sub = a - b
	sum = a + b 
	returnn // 由于对返回值命名，因此，return后面可以省略
}

func main() {
	ret1, ret2 := test(10, 20)
	fmt.Println(ret1, ret2) // 30 -10
}
	