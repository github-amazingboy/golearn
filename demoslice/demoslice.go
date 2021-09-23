package main

import "fmt"

func main() {
	var a = []int{1, 2, 3, 4, 5}
	a = append(a, 6)
	fmt.Println(a)
	// 要删除索引为2的元素
	a = append(a[:2], a[3:]...)

	fmt.Println(a)
}
