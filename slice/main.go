package main

import "fmt"

func main() {
	fmt.Println("切片相关知识")
	fmt.Println("声明切片的格式是： var name []type（不需要说明长度） 未初始化之前默认为 nil，长度为 0  计算容量的函数 cap()")

	nums := []int{1, 2, 3}
	fmt.Printf("相应值的go语法表示%#v  相应值 %v\n", nums, nums)

	var arr1 [6]int
	var slice1 []int = arr1[2:5] // 截取数组的地址  数组的切片也是切片类型
	arr2 := [5]int{1, 2, 3, 4, 5}
	slice2 := arr2[1:]
	fmt.Println("len和 cap都一样", arr2, slice2, len(arr2), len(slice2), cap(slice2))

	// 给数组赋值
	for i := 0; i < len(arr1); i++ {
		arr1[i] = i
	}

	for i := 0; i < len(slice1); i++ {
		fmt.Printf("Slice 第 %d位是 %d\n", i, slice1[i])
	}
	slice3 := make([]int, 3, 10)
	fmt.Printf("make函数构造切片 %T %d %d\n", slice3, len(slice3), cap(slice3))

	fmt.Printf("arr的长度是 %d\n", len(arr1))
	fmt.Printf("slice1的长度是 %d\n", len(slice1))
	// 这下面可能不太好理解  就是 切片的容量  =  数组/切片的 长度/容量 - 左指针走过的数（也就是start）
	fmt.Printf("slice1的容量是 %d  切片的长度永远不会超过它的容量   0 <= len(s) <= cap(s)  \n", cap(slice1))
	fmt.Println("因为容量 > len，所以可以调用append方法增加长度 append 返回一个 切片", append(slice1, 10))
	fmt.Println("长度没有增加，因为没有赋值", slice2, len(slice1), cap(slice1))
	fmt.Println("如果多次append 多次重新赋值 超出的本身容量的上线，会根据一个算法自动扩容 1-1.5倍之间")

	// 副作用
	fmt.Println("切片之间公用了一个地址，所以在 容量范围内  一个切片的变化会影响到他的切片。 数组也如此。。    结束是 当超过容量本身之后，回去重新申请空间。所以这个副作用就消失了")

	// copy函数 复制切片  长度取决于 复制的长度。超出长度就不复制了

	slice4 := []int{1, 2, 3}
	slice5 := []int{10, 20, 30, 40, 50}

	copy(slice4, slice5)
	fmt.Println(slice4) //[10 20 30]

	slice4 = []int{1, 2, 3}
	copy(slice5, slice4)
	fmt.Println(slice5) //[1 2 3 40 50]

	// 删除
}
