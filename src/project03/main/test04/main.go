package main

import (
	"fmt"
	"math/rand"
	"time"
)
func main()  {
	//我们为了生成一个随机数，还需要给rand设置一个种子
	//如何随机生成1-100的数数
	// rand.Seed(time.Now().Unix())
	// n := rand.Intn(100) + 1 //[0,100)
	// fmt.Println("n=",n)


	var count int = 0
	for {
		rand.Seed(time.Now().UnixNano())
		n := rand.Intn(100) + 1
		fmt.Println("n=",n)
		count ++
		if n == 99 {
			break//跳出for循环
		}
	}
	fmt.Println("生成99一共用了",count)
}