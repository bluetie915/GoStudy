package main

// import "fmt"
// import "unsafe"
import(
	"fmt"
	"unsafe"
)

//整型变量学习
func main(){
	//有符号整数的各个类型参考图int
	//无符号整数的各个类型参考图uint
	//特殊的整数的各个类型参考图other
	var i int8 = 10
	fmt.Println("i=",i)

	var b uint = 111
	fmt.Println("b=",b)
	var c byte = 255
	fmt.Println("c=",c)

	//类型使用的细节
	var x = 100		//x是什么数据类型
	fmt.Printf("x的类型是%T\n",x)

	//如何在程序查看某个变量的占用字节大小和数据类型
	var y int64 = 10
	fmt.Printf("y的类型为%T,y占用的字节数是%d",y,unsafe.Sizeof(y))

	//golang程序中整型变量在使用时，遵守保小不会保大的原则
	//即：在保证程序正确运行下，尽量使用占用空间小的数据类型
	var age byte = 45
	fmt.Println("年龄方便使用这个",age)
}