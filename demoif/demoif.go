package main

import "fmt"

func ifDemo1() {
	score := 65
	if score >= 90 {
		fmt.Println("A")
	} else if score > 75 {
		fmt.Println("B")
	} else {
		fmt.Println("C")
	}
}

func ifDemo2() {
	if score := 65; score >= 90 {
		fmt.Println("A")
	} else if score > 75 {
		fmt.Println("B")
	} else {
		fmt.Println("C")
	}
}

func forDemo() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

func forRange() {
	var a = []int{1, 2, 3, 4, 5}
	for i, v := range a {
		fmt.Printf("i is %d,v is %d \n", i, v)
	}

	var b = map[int]string{1: "helo", 2: "world"}
	for key, value := range b {
		fmt.Printf("key is %d, v is %s \n", key, value)
	}

	var c = make(map[string]string, 3)
	c["tom"] = "cat"
	c["jack"] = "mouse"
	for mk, mv := range c {
		fmt.Printf("mk is %s , mv is %s \n", mk, mv)
	}
}

func main() {
	ifDemo1()
	ifDemo2()
	forDemo()
	forRange()
}
