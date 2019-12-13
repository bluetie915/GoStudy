package main

import "fmt"

//定义全局变量
// var a1 = 100
// var a2 = 200
// var a3 = "jack"

//上面的声明方式也可以改成一次性声明
var(
	a1 = 100
	a2 = 200
	a3 = "jack"
)
func main(){
	//演示golang如何一次性声明多个变量
	var n1,n2,n3 int
	fmt.Println("n1=",n1,"n2=",n2,"n3=",n3)

	//一次性声明多个变量的方式2
	var m1, name, m3 = 100, "tom", 888
	fmt.Println("m1=",m1,"name=",name,"m3=",m3)

	//一次性声明多个变量的方式3，同样可以使用类型推导
	x1, firstName, x3 := 100, "tom", 888
	fmt.Println("x1=",x1,"firstName=",firstName,"x3=",x3)

	//输出全局变量
	fmt.Println("a1=",a1,"a2=",a2,"a3=",a3)
}