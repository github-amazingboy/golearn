package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Fprintln(os.Stdout, "向标准输出写入内容")
	fileObj, err := os.OpenFile("./xx.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("打开文件出错,err:", err)
		return
	}
	name := "沙河小王子"
	fmt.Fprintf(fileObj, "往文件中写信息:%s", name)
}
