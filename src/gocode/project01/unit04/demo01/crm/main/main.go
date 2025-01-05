package main //1、package进行包的声明，建议：包名和目录名保持一致
//2、main包是程序的入口包，一个程序只能有一个main包,一般main函数都在main包下

//3、包名都是从$GOPATH/src后开始的，使用/进行路径分割。
import (
	"fmt"
	test "gocode/project01/unit04/demo01/crm/calutils" //3、导入aaa包，aaa也可以是其他别名test，起了别名后原来的包名就不能使用了
	"gocode/project01/unit04/demo01/crm/dbutils"       //3、导入dbutils包
)

// 4、main函数是程序的入口函数，程序的执行从main函数开始
func main() {
	fmt.Println("Hello, World!")
	dbutils.GetConn() //5、调用dbutils包的GetConnection函数
	test.Add()        //
}
