package	main
import (
	"fmt"
)
//通过指针改变指向值
func main() {
	var a int = 10
	var b *int = &a
	fmt.Println(a, *b) // 10 10	
	a = 20
	fmt.Println(a, *b) // 20 20
	*b = 30  // 通过指针改变指向值
	fmt.Println(a, *b) 	// 30 30
}

// 指针接收的一定是个变量的地址，而不是常量或者表达式的值
// 通过指针改变指向值，只需要在指针前加*即可，如*b = 30
// 基本数据类型都有对应的指针类型，如int对应*int，string对应*string
// 指针类型的零值是nil，即空指针