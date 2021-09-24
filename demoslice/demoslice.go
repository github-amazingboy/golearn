package main

import (
	"fmt"
	"sort"
)

func main() {
	var a = []int{1, 2, 3, 4, 5}
	a = append(a, 6)
	fmt.Println(a)
	// 要删除索引为2的元素
	a = append(a[:2], a[3:]...)
	fmt.Println(a)

	test2()
	test3()
}
func test2() {
	var a = make([]int, 5, 10)
	for i := 0; i < 10; i++ {
		a = append(a, i)
	}
	fmt.Println(a)
}
func test3() {
	var a = []int{3, 7, 8, 9, 1}
	//sort.Sort(sort.IntSlice(a))
	sort.Ints(a)
	fmt.Println(a)
}
