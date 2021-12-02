package main

import "fmt"

func main() {

	slice := [][]int{}
	slice2 := make([][]int, 0)
	fmt.Println(slice, slice2)

	// 没啥学的  基本原理相通
	slice = append(slice, []int{1, 2, 3})

	fmt.Println(slice, slice2)

}
