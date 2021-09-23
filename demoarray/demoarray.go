package main

import "fmt"

func main() {
	var testArray [3]int
	var numArray = [...]int{1, 2}
	var cityArray = [...]string{"北京", "上海", "深圳"}
	fmt.Println(testArray)
	fmt.Println(numArray)
	fmt.Println(cityArray)
	fmt.Printf("type of cityArray:%T\n", cityArray)

	a := [...]int{1: 1, 3: 5}
	fmt.Println(a)
	fmt.Printf("type of a:%T\n", a)

	for i := 0; i < len(cityArray); i++ {
		fmt.Println(cityArray[i])
	}
	for index, value := range cityArray {
		fmt.Println(index, value)
	}

	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %d \t", i, j, i*j)
		}
		fmt.Println("")
	}

	erwei := [3][2]string{
		{"北京", "上海"},
		{"广州", "深圳"},
		{"成都", "重庆"},
	}
	fmt.Println(erwei)
	fmt.Println(erwei[2][1])

	for _, v1 := range erwei {
		for _, v2 := range v1 {
			fmt.Println(v2)
		}
	}
	t1 := [3]int{1, 2, 3}
	modifyArray(t1)
	fmt.Printf("t1 %v", t1)
}

func modifyArray(x [3]int) {
	x[0] = 100
}
func modifyArray2(x [3][2]int) {
	x[2][0] = 100
}
