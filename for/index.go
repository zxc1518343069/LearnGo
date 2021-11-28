package main

import "fmt"

func main() {
	// For 类似 py

	count := 0
	for i := 0; i < 100; i++ {
		count += i
	}
	fmt.Println("也有continue   break")
	fmt.Println("break 能结合 goto 使用 不过这会破坏一些规则 所以不建议")

	// 类似while
	//for i < 100 {
	//
	//}

	// 死循环
	//for {
	//
	//}
	fmt.Printf("总和为 %d", count)

	str := "叫我阿琛"

	for index, item := range str {
		fmt.Printf("%d %T %c %T \n", index, index, item, item)
	}
}
