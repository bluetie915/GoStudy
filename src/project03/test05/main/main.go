package main

import (
	"fmt"
	 util "project03/test05/utils"
)
//包的使用
func main()  {
	var result = util.Add(3,5)
	fmt.Println("result=",result)

	//跨包必须大写，小写为私有
	var flag int = util.Num
	fmt.Println("flag=",flag)
}