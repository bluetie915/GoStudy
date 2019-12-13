package main

import "fmt"

func main(){
	//golang的变量使用

	//1>指定变量类型，声明数据类型，若不赋值，使用默认
	var i int
	fmt.Println("i=",i)

	//2>根据值自行判定变量类型(类型推导)
	var num = 10.11
	fmt.Println("num=",num)

	//3>省略var,注意 := 左侧的变量不应该是已经声明过的，否则会导致编译错误
	//等价var name string   name = "tom"
	//其中的:不能省略
	name := "tom"
	fmt.Println("tom=",name)
}