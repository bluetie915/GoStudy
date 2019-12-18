package main

import (
	"fmt"
)
func main(){
	var age byte
	fmt.Println("请输入你的年龄：")
	fmt.Scanln(&age)
	if age >18 {
		fmt.Println("恭喜你已经成年了")
	}
}