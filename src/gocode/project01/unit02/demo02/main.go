package main
import "fmt"
 
//声明全局变量(写在函数外的变量),不能使用 := 的方式;只能使用 var() 的方式,或者直接赋值,不能省略 var,否则会报错
var q1 = 100
//name := "tom" //这种写法错误，不能使用
var (
	q2 = 200
	q3 = "jamce"
)


func main(){
	//1.变量的声明(写在函数内的变量)
	//1、第一种：变量的使用方式：指定变量类型，并且赋值
	var age int = 18
	fmt.Println("age=", age)

	//2、指定变量的类型，但是不赋值。使用默认值
	var num int
	fmt.Println("num=", num)

	//3、第三种：省略var，注意 := 左侧的变量不应该是已经声明过的，否则会导致编译错误
	// 下面的方式等价 var name string    
	name2 := "tom"

	//4、一次性声明多个变量
	var n1, n2, n3 int
	fmt.Println("n1=", n1, "n2=", n2, "n3=", n3)

	//5、一次性声明多个变量的方式2
	var n4, name3, n5 = 100, "tom", 888
	fmt.Println("n4=", n4, "name3=", name3, "n5=", n5)

	//6、一次性声明多个变量的方式3，可以类型推导
	n6, name4, n7 := 100, "tom~", 888
	fmt.Println("n6=", n6, "name4=", name4, "n7=", n7)

	//7、输出全局变量
	fmt.Println("q1=", q1, "q2=", q2, "q3=", q3)
}

