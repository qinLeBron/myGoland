package	main
import "fmt"

// 定义一个函数
	// 函数的定义
	// func 函数名(入参数列表) (返回值类型) {
	//		函数体
	// }

// 定义一个函数，计算两个数的和
func cal (num1 int, num2 int) int { // num1, num2 是形参   返回值类型:如果只有一个的话，可以省略小括号；如果没有返回值，也可以省略小括号
	return num1 + num2
}

//定义函数，计算两个数的和 和 差
func cal2 (num1 int, num2 int) (int, int) {
	sum := num1 + num2
	sub := num1 - num2
	return sum, sub
}


func main() {
	// 调用函数
	// 函数名(实参列表)
	result := cal(10, 20)
	fmt.Println("result = ", result)
	// 调用函数

	sum, sub := cal2(10, 20) // sum, sub 是接收返回值的变量  不想接收的话，可以使用 _ 占位符
	fmt.Println("sum = ", sum)
	fmt.Println("sub = ", sub)

}