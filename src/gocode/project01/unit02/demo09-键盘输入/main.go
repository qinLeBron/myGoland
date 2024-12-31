/*
键盘输入有两个函数可以使用
fmt.Scanln() 输入完毕自动换行
fmt.Scanf() 输入完毕不会自动换行
接收参数为指针类型，即&name
*/
package main
import "fmt"
func main() {
	var name string
	var age byte
	var sal float32
	var isPass bool
	// fmt.Println("请输入姓名：")	// fmt.Scanln() 会在输入完毕后自动换行
	// fmt.Scanln(&name) 	
	// fmt.Println("请输入年龄：")
	// fmt.Scanln(&age)
	// fmt.Println("请输入薪水：")
	// fmt.Scanln(&sal)
	// fmt.Println("请输入是否通过：")
	// fmt.Scanln(&isPass)
	fmt.Println("请输入姓名 年龄 薪水 是否通过：")
	fmt.Scanf("%s %d %f %t", &name, &age, &sal, &isPass) // fmt.Scanf() 不会在输入完毕后自动换行
	fmt.Printf("名字是 %v\n年龄是 %v\n薪水是 %v\n是否通过 %v\n", name, age, sal, isPass)
}