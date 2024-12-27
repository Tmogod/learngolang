package util

import "fmt"

var (
	Name = "二郎显圣真君"
)

func Add(a, b int) int {
	return a + b
}

func init() {
	fmt.Printf("init util\n")
}
