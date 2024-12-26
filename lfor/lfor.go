package main

import "fmt"

func main() {
	//for i := 0; i < 3; i++ {
	//	fmt.Println(i)
	//}

	//var i int = 0
	//for ; i < 3; i++ {
	//	fmt.Println(i)
	//}

	//i := 0
	//for ;i<3; {
	//	fmt.Println(i)
	//	i++
	//}

	i := 0
	for i < 3 {
		fmt.Println(i)
		i++
	}
	// 死循环
	for {
		if i > 10 {
			break // break 退出for循环
		} else if i < 10 {
			continue // 跳过此次循环，不执行continue下面的代码
		}
	}
}
