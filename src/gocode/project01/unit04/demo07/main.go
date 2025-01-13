package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Printf("%v ~~~ 对应的类型是：%T\n", now, now)

	//调用结构体中的方法
	fmt.Printf("年：%v \n", now.Year())
	fmt.Printf("月：%v \n", now.Month())
	fmt.Printf("月：%v \n", int(now.Month()))
	fmt.Printf("日：%v \n", now.Day())
	fmt.Printf("时：%v \n", now.Hour())
	fmt.Printf("分：%v \n", now.Minute())
	fmt.Printf("秒：%v \n", now.Second())

	//Print将字符串直接输出
	fmt.Printf("当前时间为： %d-%d-%d %d:%d:%d \n", now.Year(), now.Month(),
		now.Day(), now.Hour(), now.Minute(), now.Second())

	//Sprintf可以得到这个字符串，以便后续使用
	datestr := fmt.Sprintf("当前时间为： %d-%d-%d %d:%d:%d \n", now.Year(), now.Month(),
		now.Day(), now.Hour(), now.Minute(), now.Second())
	fmt.Printf(datestr)

	//指定格式输出
	//这个参数字符串的各个数字必须是固定的，必须这样写
	datestr2 := now.Format("2006/01/02 15/04/05")
	fmt.Println(datestr2)

	//任意组合都是可以的，根据需求自己选择（自己任意组合）
	datestr3 := now.Format("02 15:04")
	fmt.Println(datestr3)
}
