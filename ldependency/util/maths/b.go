package maths

import "fmt"

func Add(a, b, c int) int {
	return a + sub(b, c)
}

func init() {
	fmt.Printf("math init\n")
}