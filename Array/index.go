package main

import "fmt"

func main() {
	//Go 语言中的数组是一种 值类型  可以new
	var nums [10]int
	var nums2 = [10]int{10, 20, 7: 70} // 初始化  固定位置赋值
	var nums3 = [...]int{10, 20, 30}   // 自动判断大小

	nums4 := [3]int{1, 2, 3} // 数组常量 也成为数组初始化
	nums5 := [3]int{1, 2, 4}
	nums6 := [3]int{1, 2, 3}
	fmt.Println(nums, nums[1], nums2, nums3, nums4)

	fmt.Println("可以直接数组比较？ 离谱")
	fmt.Println(nums4 == nums5)
	fmt.Println(nums4 == nums6)
	fmt.Println(&nums4 == &nums6)
	fmt.Printf(" %p  %p\n", &nums6, &nums4)
	fmt.Println(len(nums4))

	// 便利
	for i := 0; i < len(nums4); i++ {
		fmt.Println(nums4[i])
	}

	for index, value := range nums4 {
		fmt.Println(index, value)
	}

	fmt.Println("切片出来是个切片类型")
	fmt.Printf("%T", nums2[0:2])      //[]int  不是数组 数组有长度
	fmt.Printf("%v \n", nums2[0:2:3]) // 3是 指的是容量

	var marrays [3][2]int
	marrays2 := [3][2]int{{1, 2}, {2, 3}, {3, 4}}
	fmt.Println("多维数组", marrays, marrays2)

	//
	fmt.Println("Go 语言中的数组是一种 值类型  可以new")
	var arr1 = new([5]int) // arr1应该是一个指针1
	arr2 := *arr1
	arr2[2] = 100
	fmt.Printf("%T %T\n", arr1, arr2)           //*[5]int [5]int
	fmt.Printf("%p %p %p\n", arr1, arr2, &arr2) //*[5]int [5]int
	fmt.Println(arr1, arr2)
}
