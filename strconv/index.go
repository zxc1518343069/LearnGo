package main

import (
	"fmt"
	"strconv"
)

// https://studygolang.com/pkgdoc 需要自查
func main() {
	fmt.Println("字符串转换")
	nums1 := 1
	str := fmt.Sprintf("%d", nums1) // 其他类型转换成字符串类型
	parseInt, err := strconv.ParseInt("11", 2, 64)
	if err != nil {
		return
	}
	fmt.Println(parseInt)

	fmt.Printf("%T %s", str, str)
}
