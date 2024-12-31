/*
switch语句
	1、switch后面是个表达式，这个表达式的值和case后面的表达式的值做比较
	2、如果匹配成功，则执行对应的case语句块，执行完毕后退出switch语句
	3、case后面的表达式的数据类型，必须和switch后面的表达式数据类型一致
	4、case后面的表达式可以是常量值、变量值、或者表达式
	5、多个case后面的表达式的值不能重复
	6、default语句是可选的，可以出现在任何位置，也可以没有
	7、fallthrough语句可以穿透case，继续执行下一个case，但是不建议使用
	8、switch后面的表达式可以省略，省略之后，相当于直接作用true
	9、switch后面的表达式可以是一个变量或者一个表达式
	10、case后面可以跟多个表达式，使用逗号分隔

*/

package main
import (
	"fmt"
)

func main() {
	// switch 后面的表达式
	//var num int = 5
	switch num := 5; num {  //
	case 1:
		fmt.Println("按下的是1楼")
	case 2:
		fmt.Println("按下的是2楼")
	case 3:
		fmt.Println("按下的是3楼")
	case 4:
		fmt.Println("按下的是4楼")
	default:
		fmt.Println("按下的是xxx楼")
	}

	// case后面的表达式
	var score int = 85
	switch {
	case score > 90:
		fmt.Println("成绩优秀")
	case score > 80:
		fmt.Println("成绩良好")
	case score > 70:
		fmt.Println("成绩一般")
	default:
		fmt.Println("成绩不合格")
	}

	// case后面的表达式的值不能重复
	// var s int = 1
	// switch s {
	// case 1:
	// 	fmt.Println("按下的是1楼")
	// case 1:
	// 	fmt.Println("按下的是1楼")
	// }

	// fallthrough语句
	var n int = 10
	switch n {
	case 10:
		fmt.Println("ok1")
		fallthrough // 穿透 可以继续执行下一个case  不建议使用
	case 20:
		fmt.Println("ok2")
	case 30:
		fmt.Println("ok3")
	default:
		fmt.Println("没有匹配到")
	}
}