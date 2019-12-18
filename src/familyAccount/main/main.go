package main

import (
	"fmt"
	"familyAccount/utils"
)
func main()  {
	fmt.Println("这个是面向对象的方式完成的...")
	utils.NewFamilyAccount().MainMenu()
}