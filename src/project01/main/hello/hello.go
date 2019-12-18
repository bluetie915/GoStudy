package main
//学习测试helloworld
/*
	go build 也可自定义生成的.ext的名字
	go build -o myhello.exe hello.go
	1>.go命名文件
	2>go应该程序的执行入口是main()方法
	3>go语言严格区分大小写
	4>go不需要在后面写;
	5>go中每条语句独占一行
	6>go定义的变量和import的包必须用到
	7>go中的大括号都是成对出现的
 */
//引入包
import "fmt"
func main(){
	fmt.Println("hello,world!");
	fmt.Println("hello,world!");
	fmt.Println("hello,world!");
	fmt.Println("hello,world!");
}