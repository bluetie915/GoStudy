package main

import (
	"fmt"
)
func main(){
	// var key byte
	// fmt.Println("请输入一个字符 a,b,c,d,e")
	// fmt.Scanf("%c",&key)

	// switch key {
	// case 'a','b':
	// 	fmt.Println("今天是周一")
	// case 'c':
	// 	fmt.Println("今天是周三")
	// case 'd':
	// 	fmt.Println("今天是周四")
	// case 'e':
	// 	fmt.Println("今天是周五")
	// default:
	// 	fmt.Println("您是输入有误")
	// }

	
	//switch的穿透fallthrough
	var num int = 10
	switch num {
	case 10:
		fmt.Println("ok1")
		fallthrough //默认只能穿透一层
	case 20:
		fmt.Println("ok2")
	case 30:	
		fmt.Println("ok3")
	default:
		fmt.Println("没有匹配到")
		
	}

}