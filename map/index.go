package main

import "fmt"

func main() {
	var maps2 map[string]int
	maps := map[string]int{"achen": 1, "cc": 2}
	fmt.Printf("%T %#v\n", maps, maps)     // map[string]int map[string]int{}
	fmt.Printf("%T %#v\n", maps2, maps2)   // map[string]int map[string]int(nil)
	fmt.Println(maps == nil, maps2 == nil) // false true

	maps["aa"] = 1
	v, ok := maps["achen"] // 返回   值和 存不存在
	fmt.Println(v, ok)     // 1 true
	v, ok = maps["achen2"] // 返回值和 存不存在
	fmt.Println(v, ok)     //0 false
	// 一种另类写法
	if v2, ok2 := maps["achen"]; ok2 {
		fmt.Println(v2)
	}
	// 删除
	delete(maps, "cc")
	delete(maps, "a2a")                               // 删除没有的也没报错
	fmt.Printf("%T %#v %d \n", maps, maps, len(maps)) // map[string]int map[string]int{"aa":1, "achen":1}

	for key, val := range maps {
		fmt.Println(key, val)
	}
}
