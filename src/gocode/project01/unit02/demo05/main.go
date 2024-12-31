package main
importt (
	"fmt"
)
//指针
//Go语言中的指针不能进行偏移和运算，是安全指针

func main() {
	//基本数据类型在内存布局
	var i int = 10
	//i 的地址是什么，&i
	fmt.Println("i的地址=", &i)

	//1. ptr 是一个指针变量
	//2. ptr 的类型 *int
	//3. ptr 本身的值&i
	//&i  i的内存空间	是ptr变量具体的值
	//*ptr 代表指针变量指向的具体数据
	var ptr *int = &i	//ptr接受的是i的地址
	fmt.Printf("ptr=%v\n", ptr)
	fmt.Printf("ptr的地址=%v\n", &ptr)
	fmt.Printf("ptr指向的值=%v\n", *ptr)

	/*
	总结：最重要的就是两个符号
	& 取内存地址
	* 根据地址取值
	*/
}