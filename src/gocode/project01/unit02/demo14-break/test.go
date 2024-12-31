//break 带标签的使用
package main
import "fmt"

func main() {
	// break语句使用
	lable1:
	for i := 1; i <= 4; i++ {
		// lable2:
		for j := 1; j <= 10; j++ {
			if j == 2 {
				// break lable1
				break lable1
			}
			fmt.Println("i=", i, "j=", j)
		}
	}
}
