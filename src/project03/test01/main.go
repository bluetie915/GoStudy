package main

import "fmt"
func main()  {
	//字符串的遍历，传统方式
	var param string = "hello,world"
	for i := 0;i < len(param); i++{
		fmt.Printf("%c",param[i])
	}

	fmt.Println()

	//字符串遍历，for-range方式 
	param = "abc-ok"
	for index, val := range param{
		fmt.Printf("index=%d,val=%c\n",index,val)
	}
}