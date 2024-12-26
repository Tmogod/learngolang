package main

import "fmt"

func main() {
	// 声明变量
	var a uint8
	var b uint16
	var c uint32
	var age byte
	var sex bool
	var price float32
	var totalPrice float64

	// 变量的赋值
	a = 10
	b = 20
	c = uint32(a) + uint32(b)

	fmt.Printf("c=%d\n", c)

	// := 声明+赋值
	f := 40.0                  // float64
	g := uint32(c) - uint32(a) // int

	var h int = 4
	var m = 4

	// 同时声明多个变量
	l, n := 20, true

	// 未显式声明的变量，其值是"0"
	fmt.Printf("age类型为: %T, age = %d\n", age, age)
	fmt.Printf("sex类型为: %T, sex = %t\n", sex, sex)
	fmt.Printf("price类型为: %T, price = %.2f\n", price, price)
	fmt.Printf("totalPrice类型为: %T, totalPrice = %.2f\n", totalPrice, totalPrice)
	fmt.Printf("f类型为: %T, f = %.2f\n", f, f)
	fmt.Printf("g类型为: %T, g = %d\n", g, g)
	fmt.Println(h)
	fmt.Println(m)
	fmt.Println(l)
	fmt.Println(n)

}
