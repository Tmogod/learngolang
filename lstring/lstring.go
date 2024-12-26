package main

import "fmt"

func main() {
	// 字符串是不可修改的byte[]数组
	s := "golang"
	fmt.Printf("%s\n", s)

	for i, e := range s {
		fmt.Printf("%d: %c\n", i, e)
	}
	fmt.Printf("string s's len: %d\n", len(s))

	a := "golang你好"            // UTF8编码
	fmt.Printf("%d\n", len(a)) // 12

	arr := []rune(a)             // rune 代表UTF8编码的每一个元素
	fmt.Printf("%d\n", len(arr)) // 8

	for _, e := range arr {
		fmt.Printf("%c\n", e) // %c能输出"你好"等编码字符
	}
	// ``会取消转义
	b := `golang\n
    	你好
	hello`
	fmt.Printf("%s\n", b)
	// 字符串的拼接
	c := "golang"
	d := "你好"
	f := c + " " + d
	fmt.Printf("%s\n", f)
}
