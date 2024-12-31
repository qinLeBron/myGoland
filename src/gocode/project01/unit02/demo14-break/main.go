/*
 break语句使用
 	break语句只能在循环和switch语句中使用
	break语句可以结束for、switch和select的代码块
	break语句后面还可以带标签，表示退出某个标签对应的代码块
break的作用是结束离它最近的循环
*/

package main
import "fmt"

func main() {
	// break语句使用
	var sum int = 0
	for i := 1; i <= 100; i++ {
		sum += i
		if sum >= 300 {
			break
		}
		fmt.Println(i)
	}
	fmt.Println("over")



	// break语句后面还可以带标签，表示退出某个标签对应的代码块
	// break的作用是结束离它最近的循环
	// break语句还可以结束指定的循环
	
}