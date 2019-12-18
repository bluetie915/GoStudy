package main

import (
	"fmt"
)
func main(){
	for i := 0; i <= 10; i++ {
		fmt.Println("恭喜你输出成功",i)
	}

	//for循环的第二种写法
	j := 1
	for j <= 10 {
		fmt.Println("输出成功",j)
		j++
	}

}