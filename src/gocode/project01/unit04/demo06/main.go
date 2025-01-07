package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	//统计字符串的长度，按字节进行统计
	str := "golang你好"     //汉字是utf-8字符集，一个汉字3个字节
	fmt.Println(len(str)) //12  会换行
	fmt.Print(str)
	fmt.Printf(str)

	//对字符串进行遍历
	//方式1：利用键值循环：for-range
	for i, value := range str {
		fmt.Printf("索引为：%d, 具体的值为：%c \n", i, value)
	}

	//方式2：利用r:=[]rune(str)
	r := []rune(str)
	for i := 0; i < len(r); i++ {
		fmt.Printf("%c \n", r[i])
	}

	//字符串转整数
	num1, _ := strconv.Atoi("555")
	fmt.Println(num1)

	//整数转字符串
	str1 := strconv.Itoa(666)
	fmt.Println(str1)

	//统计一个字符串有几个指定的的子串：
	count := strings.Count("golangandpythongo", "go")
	fmt.Println(count)

	//不区分大小写的字符串比较
	flag := strings.EqualFold("hello", "HeLLo")
	fmt.Println(flag)

	//区分大小写字符串比较：
	fmt.Println("Hello" == "HellO")

	//返回子串在字符串第一次出现的索引值，如果没有返回-1：
	fmt.Println(strings.Index("golangandpython", "gan"))

	//字符串的替换
	str11 := strings.Replace("gogoandpythongogo", "go", "golang", -1)
	str2 := strings.Replace("gogoandpythongogo", "go", "golang", 2)
	fmt.Println(str11)
	fmt.Println(str2)

	// 按照指定的某个字符，为分割标识，将一个字符串拆分成字符串数组
	arr := strings.Split("golang-python-java", "-")
	fmt.Println(arr)

	//将字符串的字母进行大小写的转换
	fmt.Println(strings.ToLower("GoGO"))
	fmt.Println(strings.ToUpper("gogo"))

	//将字符串两边的空格去掉
	fmt.Println(strings.TrimSpace("     ha ha ha     "))

	//将字符串两边的字符去掉
	fmt.Println(strings.Trim("~ha ha ha~", "~"))

	//将字符串左边指定的字符去掉
	fmt.Println(strings.TrimLeft("~ha ha ha~", "~"))

	//将字符串右边指定的字符去掉
	fmt.Println(strings.TrimRight("~ha ha ha~", "~"))

	//判断字符串是否以指定的字符开头
	fmt.Println(strings.HasPrefix("https://www.baidu.com", "https")) //true

	//判断字符串是否以指定的字符结尾
	fmt.Println(strings.HasSuffix("https://www.baidu.com", ".cn")) //false

}
