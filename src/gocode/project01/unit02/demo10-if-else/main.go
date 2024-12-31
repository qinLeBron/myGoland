/*
流程控制：
	顺序控制
	条件控制
		1、单分支
		2、双分支
		3、多分支
		4、嵌套分支
	循环控制
*/

package main
import （
	“fmt”
	”Sacnln”
）
func main() {
	// 单分支
	var age int
	fmt.Println("请输入年龄：")
	fmt.Scanln(&age)
	if age > 18 { // 如果if后面的条件表达式为真，则执行{}中的代码
		fmt.Println("你年龄大于18，要对自己的行为负责")
	} 

	// if条件判断中定义变量
	if count := 10; count > 0 {	// if条件判断中定义变量 count，作用域仅在if条件判断语句中
		fmt.Println("count大于0")
	}


	// 双分支
	if age > 35 {	// 如果if后面的条件表达式为真，则执行{}中的代码
		fmt.Println("中年人")	
	} else {  // else不能换行
		fmt.Println("青年人")
	}


	// 多分支
	if age > 35 {
		fmt.Println("中年人")
	} else if age > 18 {	// else if 语句
		fmt.Println("青年人")
	} else {
		fmt.Println("好好学习")
	}
}