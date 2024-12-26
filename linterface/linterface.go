package main

import "fmt"

// 接口是一组行为规范的集合
type Human interface {
	Say(int, int) int
}

func foo(h Human) {
	c := h.Say(3, 6)
	fmt.Println(c)
}

func main() {
	var a Human

	t := Tom{}
	a = t

	foo(a)

	j := Jim{}
	a = j

	foo(a)

	var b interface{} // 空接口类型，匹配任何对象
	var c []string
	b = c

	fmt.Printf("%v\n", b)
}

type Tom struct {
}

func (t Tom) Say(i, j int) int {
	fmt.Printf("Say %d %d \n", i, j)
	return i - j
}

type Jim struct {
}

func (t Jim) Say(i, j int) int {
	fmt.Printf("Say %d %d \n", i, j)
	return i + j
}
