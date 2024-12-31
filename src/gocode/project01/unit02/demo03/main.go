package main
import (
	"fmt"
	"strconv"
)
//基本数据类型转换为string
func main() {
	var num1 int = 99
	var num2 float64 = 23.456
	var b bool = true
	var myChar byte = 'h'
	var str string
	//使用第一种方式 fmt.Sprintf方法 
	//推荐使用Sprintf方法
	str = fmt.Sprintf("%d", num1)  //%d 整型格式 10进制
	fmt.Printf("str type %T str=%q\n", str, str) //%T 打印数据类型 %q 打印字符串
	str = fmt.Sprintf("%f", num2) //%f 浮点型格式
	fmt.Printf("str type %T str=%q\n", str, str)
	str = fmt.Sprintf("%t", b) //%t 布尔型格式
	fmt.Printf("str type %T str=%q\n", str, str)
	str = fmt.Sprintf("%c", myChar) //%c 字符格式
	fmt.Printf("str type %T str=%q\n", str, str)
	//第二种方式 strconv 函数
	var num3 int = 99
	var num4 float64 = 23.456
	var b2 bool = true
	var myChar2 byte = 'h'
	str = strconv.FormatInt(int64(num3), 10) //参数1表示要转换的数，并且必须是int64类型，参数2表示要转换的进制
	fmt.Printf("str type %T str=%q\n", str, str)
	str = strconv.FormatFloat(num4, 'f', 10, 64) //f格式，10表示小数位保留10位，64表示这个小数是float64
	fmt.Printf("str type %T str=%q\n", str, str)
	str = strconv.FormatBool(b2)
	fmt.Printf("str type %T str=%q\n", str, str)
	str = strconv.FormatUint(uint64(myChar2), 10)
	fmt.Printf("str type %T str=%q\n", str, str)
	//第三种方式 strconv 函数
	var num5 int = 99
	var num6 float64 = 23.456
	var b3 bool = true
	var myChar3 byte = 'h'
	str = strconv.Itoa(num5)
	fmt.Printf("str type %T str=%q\n", str, str)
	str = strconv.FormatFloat(num6, 'f', 10, 64) //f格式，10表示小数位保留10位，64表示这个小数是float64
	fmt.Printf("str type %T str=%q\n", str, str)
	str = strconv.FormatBool(b3)
	fmt.Printf("str type %T str=%q\n", str, str)
	str = strconv.Itoa(int(myChar3))
	fmt.Printf("str type %T str=%q\n", str, str)

}