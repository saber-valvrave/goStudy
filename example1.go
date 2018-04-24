package main

import (
	"fmt"
)

const (
	num = iota //常量生成器，每次调用自增1
	num1
	num2
)

func main() {
	fmt.Println("vim-go")
	/*
		fmt.Println(RePoint())
		var p *int = RePoint()
		*p = 2
		fmt.Printf("*p = %d\n", *p)
	*/

	//NumPrintf()
	//fmt.Printf("%d %d %d\n", num, num1, num2)
	//ArrayTest()
	//SliceTest()
	//TestStack()
	//MapTest()
	//StructTest()
	//ErrorPrint()
	//fmt.Println(ArgFunc(1,2,3,4,5,6))
	//HandleCore()
	MethodTest()
}
