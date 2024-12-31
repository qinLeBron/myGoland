/*
range 用于迭代数组(array)、切片(slice)、通道(channel)或集合(map)的元素。在数组和切片中它返回元素的索引值，在集合中返回 key-value 对的 key 值。
	1、数组、切片、字符串返回索引和值
	2、map返回key和value
	3、通道只返回通道内的值
	格式为：
		for key, value := range oldMap {
			newMap[key] = value
		}
	
*/
package main
import "fmt"

func main() {
	// 1、数组、切片、字符串返回索引和值
	nums := []int{2, 3, 4} // 切片
	sum := 0
	for i, num := range nums {
		sum += num
		fmt.Println("index:", i, "num:", num)
	}
	fmt.Println("sum:", sum)

	// 2、map返回key和value
	kvs := map[string]string{"a": "apple", "b": "banana"} // map
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	// 3、通道只返回通道内的值
	c := make(chan int)
	go func() {
		c <- 1
		c <- 2
		c <- 3
		close(c)
	}()
	for v := range c {
		fmt.Println(v)
	}
}