package main

import "fmt"

//演示golang小数的使用
func main(){
	var price float32 = 434.43
	fmt.Println("价格为",price)
	//浮点数默认都有符号
	var num1 float32 = 3434645.343
	var num2 float64 = -434545.56677
	fmt.Println("num1=",num1,"\tnum2=",num2)

	//可能造成精度损失
	var num3 float32 = -123.0000901
	var num4 float64 = -123.0000901
	fmt.Println("num3=",num3,"\tnum4=",num4)
	//float32会造成精度损失

	//golang的浮点数类型默认是float64类型
	var num5 = 4.3
	fmt.Printf("num5的数据类型是%T\n",num5)

	//  .1233 = 0.1233

	//支持科学计数法
	num8 := 4.347676e2
	fmt.Println("num8=",num8)


}
