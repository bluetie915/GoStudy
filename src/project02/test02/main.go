package main

import (
	"fmt"
)
//演示golang中string类型的使用
func main(){
	//string的基本使用
	var address string = "天津市西青区天津城建大学"
	fmt.Println(address)

	// var str = "hello"
	// str[0] = 'a'//这里不能修改str的内容，即go中的字符串不可变

	var str2 = `
		package main

		import (
			"fmt"
			"unsafe"
		)
		func main(){
			var b = false
			fmt.Println("b=",b)
			//注意事项
			//1>bool类型占用存储空间是1个字节
			fmt.Println("b的占用空间是=",unsafe.Sizeof(b))
			//2>bool类型只能取true或者false
		}
	`
	fmt.Println(str2)

	//注意事项和使用细节
	//1>go语言的字符串的字节使用utf-8
	//2>字符串一旦赋值，不可变
	//3>字符串两种表示形式分别是  双引号和反引号
		//常规双引号需要转义和特殊字符处理
		//反引号不用考虑这些问题
	//4>字符串的拼接方式
}