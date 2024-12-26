package main

import "fmt"

func main() {
	/*
		type Slice struct {
			array unsafe.Pointer
			len int
			cap int
		}
	*/

	aSlice := make([]int, 3, 5) // 创建slice(数组切片)  长度为3  容量为5
	aSlice[0], aSlice[1], aSlice[2] = 2, 7, 5

	bSlice := aSlice
	bSlice[0] = 4
	fmt.Printf("%v \n", aSlice)

	bSlice = append(aSlice, 8)  // append会构造一个新切片，容量不足会扩容， aSlice 和 bSlice不再是同一个Slice结构体, bSlice在另一个内存空间
	fmt.Printf("%v \n", aSlice) // [4 7 5]
	fmt.Printf("%v \n", bSlice) // [4 7 5 8]
	bSlice = append(bSlice, 8)
	bSlice = append(bSlice, 8)
	aSlice[0] = 1
	fmt.Printf("%v \n", aSlice) // [1 7 5]
	fmt.Printf("%v \n", bSlice) // [4 7 5 8 8 8]
	// 遍历切片
	for i, e := range bSlice {
		fmt.Printf("%d %d\n", i, e)
	}
}
