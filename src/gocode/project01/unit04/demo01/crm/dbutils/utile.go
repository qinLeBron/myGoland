// package bbb  同级别的源文件的包的声明必顿一致，
// 比如：dbutils包和utiles包的声明必须一致 package dbutils或 package bbb
package dbutils

import "fmt"

func GetConn2() { //首字母大写，可以被其他包调用
	fmt.Println("Get Connection")
}
