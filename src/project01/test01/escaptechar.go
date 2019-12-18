package main

//gofmt -w escaptechar.go能格式化代码
import "fmt"

func main() {
	fmt.Println("tom\tjack")
	fmt.Println("tom\njack")
	fmt.Println("E:\\Code\\Go\\GoStudy\\go_code\\project01\\main\\escaptechar")
	fmt.Println("tom says:\"today is a good day\"")
	// \r回车，表示从当前行的最前面开始输出，覆盖掉
	fmt.Println("我们都是好孩子，天真善良的\r孩子")
	fmt.Println("姓名\t年龄\t籍贯\t住址\njohn\t12\t河北\t北京")
}
