package main

import "fmt"

func main() {
	str()
	str2()
	var a = 1
	modify1(a)
	fmt.Println("a modify1 is ", a)
	modify2(&a)
	fmt.Println("a modify2 is ", a)
	newMake()
}

func str() {
	var s1 = "zs"
	var s2 = s1
	s1 = "ls"
	fmt.Println("string 类型打印", s1, s2)
}

func str2() {
	var s1 = "ww"
	var s2 = &s1
	*s2 = "根据地址修改"
	fmt.Printf("*s2 s1 = %v \n", s1)
}

func modify1(x int) {
	x = 5
}

func modify2(x *int) {
	*x = 6
}

func newMake() {
	var a *int
	a = new(int)
	*a = 123
	fmt.Println("a is ", *a)

	var c = make(map[string]int)
	c["是是是"] = 3
	var d = make([]int, 0)
	d = append(d, 100)
	fmt.Printf("%#v\n", c)
	fmt.Printf("%#v\n", d)
}
