package main

import (
	"fmt"
)

// 变量声明
var a = 1  // 单变量声明
var aa int // 全局变量  给定类型之后自己有一个初始值

var b, c = 2, 3 // 多变量声明 自动推导类型

var d, e int = 4, 5 // 多变量类型声明

var (
	sex = "男" // 推导类型
)

var ( // 多类型声明另一方式
	name string = "achen"
	age  int    = 18
)

var ab = a + b // 可以是计算表达式

const aaa = 1

const (
	bb int = 1
	cc
	dd
	ee        // 这些同 bb 都是1
	ff string = "123"
	gg        // 同 ff
)

const (
	zero = iota
	one
	two
)

func main() { // {}表示作用域
	// 简短声明 只能用在函数内部
	is := 1
	fmt.Println(aa, is, ab)

	// 变量交换 类似py
	b, c = c, b
	fmt.Println(b, c)

	// 打印常量
	fmt.Println("打印常量")
	fmt.Println(zero, one, two)
	fmt.Println(bb, cc, gg)
	fmt.Println("使用占位符(字符填充)打印变量")
	fmt.Printf("%T %s %d %T", is, ff, c, d) // %T 类型
}
