package main

import "fmt"
func main() {
	//go语言中没有while和do...while循环，所以要实现可以用for代替
	//例题：输出10句“hello,world”
	var i int = 1
	for {
		if i > 10 {
			break
		}
		fmt.Println("hello,world",i)
		i ++
	}
}