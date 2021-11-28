package main

import "fmt"

func main() {
	var (
		name    string
		age     int
		married bool
	)
	fmt.Println("赋值错误就显示 基本值  传的都是指针")
	//fmt.Scan( &name,&age, &married) // 将以空白符分隔的数据分别存入指定的参数。

	//fmt.Printf(" name:%s age:%d married:%t \n", name, age, married)

	//fmt.Scanf("1:%s 2:%d 3:%t", &name, &age, &married) //按照具体 格式才可以赋值成功

	fmt.Scanln(&name, &age, &married) // 遇到回车表示输入结束
	fmt.Printf(" name:%s age:%d married:%t \n", name, age, married)

}
