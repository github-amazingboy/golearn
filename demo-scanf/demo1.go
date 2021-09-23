package main

import "fmt"

func main() {
	fmt.Println("please input a string")
	var name string
	var age int
	var sal float32
	var sex bool

	_, err := fmt.Scanln(name)

	if err != nil {
		fmt.Printf("error input err=%v", err)
		return
	}

	fmt.Printf("name is %v \n", name)
	n, err := fmt.Scanln(&age)

	if err != nil {
		fmt.Print(n)
		fmt.Printf("error input err=%v", err)
		return
	}
	fmt.Scanln(&sex)

	fmt.Println("info below:")

	fmt.Printf("%s %d %f %t", name, age, sal, sex)
}
