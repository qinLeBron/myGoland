/*
for循环的基本格式如下：
for 初始表达式;条件表达式（条件判断）;结束语句{
	循环体语句
}
其中：	初始表达式：在第一次循环前执行
			初始表达式不能用var定义变量，要用:=定义
		条件表达式（条件判断）：每次循环开始前执行
		结束语句：每次循环结束时执行
*/

package main
import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}