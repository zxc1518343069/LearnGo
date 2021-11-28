package main

import "fmt"

func main() {
	A := 2
	// 指针
	var cc *int = &A //cc 指向 A的地址

	fmt.Printf("%T %d  %d\n", cc, cc, *cc)

	*cc = 3 // *cc 代表的是值
	fmt.Printf(" %d", *cc)

	// 打印指针
	fmt.Printf(" %p\n", cc)

	// %v 通用
	fmt.Printf(" %v", A)
}
