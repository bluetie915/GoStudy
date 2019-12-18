package main

import (
	"fmt"
)
func test(m int, n int) int {
	return m + n
}
func main(){
	var x = test(1, 4)
	fmt.Println("x=",x)
	var name string
	var age byte
	var salary float32
	var isPass bool

	//使用fmt.Println()

	// fmt.Println("请输入姓名：")
	// fmt.Scanln(&name)

	// fmt.Println("请输入年龄：")
	// fmt.Scanln(&age)

	// fmt.Println("请输入薪水：")
	// fmt.Scanln(&salary)

	// fmt.Println("请输入是否通过考试：")
	// fmt.Scanln(&isPass)

	// fmt.Printf(" 名字是%v\n 年龄是%v\n 薪水是%v\n 是否通过考试%v\n",name,age,salary,isPass)

	//使用fmt.Printf()
	fmt.Println("请输入你的姓名，年龄，薪水，是否通过考试，使用空格隔开")
	fmt.Scanf("%s %d %f %t",&name,&age,&salary,&isPass)
	fmt.Printf(" 名字是%v\n 年龄是%v\n 薪水是%v\n 是否通过考试%v\n",name,age,salary,isPass)
}