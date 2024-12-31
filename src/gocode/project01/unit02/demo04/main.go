package main
impoert (
	"fmt"
	"strconv"
)
//string转换为基本数据类型,
	//注意：在转换中，要确保 string 类型能够转换为有效的数据类型，比如我们可以把 "123" 转换为整型，但是不能把 "hello" 转换为整型
	//如果转换的时候，string 类型无法转成有效的数据，就会返回对应类型的默认值，比如转换为整型失败，默认值就是 0，bool类型默认值是false
func main() {
	//stgring-->bool
	var str string = "true"
	var b bool
	//ParseBool(str string) (value bool, err error)
	//函数原型 func ParseBool(str string) (value bool, err error)
	//函数功能：把字符串转换为bool类型
	//返回值：value bool 表示转换后的结果 err error 表示错误信息
	b, _ = strconv.ParseBool(str)   //,_表示忽略错误
	fmt.Printf("b type %T b=%v\n", b, b)

	//string-->int
	var str2 string = "1234590"
	var n1 int64
	//ParseInt(str string, base int, bitSize int) (i int64, err error)
	//函数原型 func ParseInt(str string, base int, bitSize int) (i int64, err error) //前面括号表示入参，后面括号表示出参
	//入参：参数1表示要转换的字符串，参数2表示要转换的进制为10进制，参数3表示转换的位数为int64
	//函数功能：把字符串转换为int64类型
	//返回值：i int64 表示转换后的结果 err error 表示错误信息
	n1, _ = strconv.ParseInt(str2, 10, 64)
	fmt.Printf("n1 type %T n1=%v\n", n1, n1)

	//string-->float
	var str3 string = "123.456"
	var f1 float64
	//ParseFloat(str string, bitSize int) (f float64, err error)
	//函数原型 func ParseFloat(str string, bitSize int) (f float64, err error)
	//函数功能：把字符串转换为float64类型
	//返回值：f float64 表示转换后的结果 err error 表示错误信息
	f1, _ = strconv.ParseFloat(str3, 64)
	fmt.Printf("f1 type %T f1=%v\n", f1, f1)
}