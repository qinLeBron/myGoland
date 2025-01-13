package main

import (
	"errors"
	"fmt"
)

func main() {
	err := test()
	if err != nil {
		fmt.Println("自定义错误：", err)
		panic(err) // 如果报错后，代码不想继续执行，需要中断推出，借助builtin包下的内置函数：panic
	}
	fmt.Println("上面的除法操作执行成功。。")
	fmt.Println("正常执行下面的逻辑。。")
}

func test() (err error) {
	num1 := 10
	num2 := 0
	if num2 == 0 {
		//抛出自定义错误：
		return errors.New("除数不能为0哦~~")
	} else {
		result := num1 / num2
		fmt.Println(result)
		// 如果没有错误，返回nil
		return nil
	}
}
