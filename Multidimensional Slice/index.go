package main

import "fmt"

func main() {

	slice := [][]int{}
	slice2 := make([][]int, 0)
	fmt.Println(slice, slice2)

	slice = append(slice, []int{1, 2, 3})
	slice[0][0] = 2
	slice[0] = make([]int, 2)

	fmt.Printf("%T \n", make([]int, 2)) // []int

	fmt.Println(slice)

	slice3 := [][]string{}
	slice3 = append(slice3, []string{"123", "23234"})

	fmt.Println(slice, slice2, slice3)

}
