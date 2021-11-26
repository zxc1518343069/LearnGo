package main

import (
	"fmt"
	"math/rand"
)

// 导入多个包   导入包里某一个函数

// go run 编译运行
// go run build -x //  打印过程  -o // 重命名  xxx.go // 文件名字
func main() {
	fmt.Println("hello world", rand.Intn(10))
}
