package main

import "fmt"
func main()  {
	for i := 1; i <= 9; i ++{
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%d ",j,i,i*j)			
		}
		fmt.Println()
	}

	fmt.Println()
	var flag = 10
	for i:=1; i <= flag; i ++ {
		for j := flag - i; j >= 1; j-- {
			fmt.Printf(" ")
		}
		for k := 1; k <= 2*i-1; k ++ {
			fmt.Printf("*")		
		}
		fmt.Println()
	}
}