/*
contuniue的使用
	1、continue语句默认用于跳过当前循环的剩余代码，然后继续进行下一轮循环
	2、continue也可以通过加标签，跳出指定标签的循环
	3、continue语句只能用在for循环中


return的使用  //结束当前函数
	1、return语句专门用于从函数中返回，后面可以跟返回值，也可以不跟
	2、return后面不跟返回值，表示返回值为空
	3、return语句在执行时，结束当前函数
*/
package main
import “fmt”
func main() {
	// 使用continue打印1-100之间的奇数
	for i := 1; i <= 100; i++ {
		if i % 2 == 0 {
			continue
		}
		fmt.Println(i)
	}
}
