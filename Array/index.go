package main

import "fmt"

func main() {
	var nums [10]int
	var nums2 = [10]int{10, 20, 7: 70} // 初始化  固定位置赋值
	var nums3 = [...]int{10, 20, 30}   // 自动判断大小

	nums4 := [3]int{1, 2, 3}
	nums5 := [3]int{1, 2, 4}
	nums6 := [3]int{1, 2, 3}
	fmt.Println(nums, nums[1], nums2, nums3, nums4)

	fmt.Println("可以直接数组比较？ 离谱")
	fmt.Println(nums4 == nums5)
	fmt.Println(nums4 == nums6)
	fmt.Println(len(nums4))

	// 便利
	for i := 0; i < len(nums4); i++ {
		fmt.Println(nums4[i])
	}

	for index, value := range nums4 {
		fmt.Println(index, value)
	}
}
