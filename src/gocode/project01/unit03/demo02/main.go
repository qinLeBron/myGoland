package demo
import "fmt"

// 可变参数的函数
func cal3 (args ...int) int {
	sum := 0
	for i := 0; i < len(args); i++ {
		sum += args[i]
	}
	return sum
}


// 基本数据类型 和 数组 默认都是值传递，即进行值拷贝。在函数内修改，不会影响到原来的值
func modify1 (a int) {
	a = 10
}	

func modify2 (a [5]int) {
	a[0] = 10
}	


// 如果想在函数内修改，可以传递指针
func modify3 (a *int) {
	*a = 10
}

func main() {
	fmt.Println(cal3(1, 2, 3, 4, 5))
	
	a := 1
	modify1(a)
	fmt.Println(a) // 1

	b := [5]int{1, 2, 3, 4, 5}
	modify2(b)
	fmt.Println(b) // [1 2 3 4 5]

	c := 1
	modify3(&c)
	fmt.Println(c) // 10
}

