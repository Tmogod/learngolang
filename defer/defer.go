package main

import "fmt"

func foo() int {
	a, b := 3, 5
	c := a + b
	defer fmt.Println("111", c) // defer代表延时执行函数, 按注册时的倒叙执行
	fmt.Println(c)
	defer fmt.Println("222", c) // defer的注册语句的时候，c会直接被赋值 打印 222 8
	defer func() {
		fmt.Println("333", c) // defer注册匿名函数的时候，是执行时去取值 打印333 100
	}()
	c = 100
	return c
}

func main() {
	foo()
}
