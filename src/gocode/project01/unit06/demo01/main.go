package main

import "fmt"

func main() {
	//实现的功能：给出5个学生的成绩，求出成绩的总和，平均数：
	// 给出5个学生的成绩：--->数组存储：
	// 定义一个数组：
	var scores [5]int // var 数组名 [数组长度]数组类型
	//将成绩存入数组：
	scores[0] = 95
	scores[1] = 91
	scores[2] = 80
	scores[3] = 78
	scores[4] = 99
	//求和
	//定义一个变量专门接收成绩的和：
	sum := 0
	for i := 0; i < len(scores); i++ {
		sum += scores[i]
	}
	avg := sum / len(scores)
	fmt.Printf("总成绩为：%v, 平均成绩为：%v", sum, avg)
}
