package testutils

import "fmt"

var Name string
var Age int
var Sex string

func init() {
	fmt.Println("testutils init()...")
	Name = "Tom"
	Age = 18
	Sex = "男"
}
