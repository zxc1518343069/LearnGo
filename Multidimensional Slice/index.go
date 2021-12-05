package main

import (
	"fmt"
	"sort"
)

func main() {

	var a = new([]int)   //
	fmt.Printf("%T ", a) // 返回的是一个指针
	slice := [][]int{}
	slice2 := make([][]int, 0)
	fmt.Println(slice, slice2)

	slice = append(slice, []int{1, 2, 3})
	slice[0][0] = 2
	slice[0] = make([]int, 2)

	// 排序  不同类型用不同的排序方式
	sorted := []int{1, 3, 2, 5, 6, 21, 3, 2, 3, 1}

	sort.Ints(sorted)

	fmt.Println(sorted)

	fmt.Printf("%T \n", make([]int, 2)) // []int

	fmt.Println(slice)

	slice3 := [][]string{}
	slice3 = append(slice3, []string{"123", "23234"})

	fmt.Println(slice, slice2, slice3)

}
