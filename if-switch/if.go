package main

import "fmt"

func main() {

	score := 0

	fmt.Println("请输入你的成绩")
	fmt.Scanln(&score)

	fmt.Printf("%T %d\n", score, score)
	if score == 100 { // 竟然不支持简写。。。。
		fmt.Println("满分")
	}
	if score < 100 && score > 90 {
		fmt.Println("A")
	} else if score < 90 && score > 80 {
		fmt.Println("B")
	} else {
		fmt.Println("其他")
	}
	fmt.Println("switch版本")

	switch score {
	case 100:
		fmt.Println("满分")

	}

	switch {
	case score < 100 && score > 90:
		fmt.Println("A")
		fallthrough // 自动执行下一个case  且不会判断 是否满足
	case score < 90 && score > 80:
		fmt.Println("B")
	default:
		fmt.Println("其他")

	}

}
