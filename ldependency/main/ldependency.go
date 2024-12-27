package main

import (
	"fmt"
	"github.com/bytedance/sonic"
	_ "ldependency/util"
	math "ldependency/util/maths"
)

func main() {
	a, b, c := 1, 2, 3
	//fmt.Println(util.Name)
	//fmt.Println(util.Add(a, b))
	fmt.Println(math.Add(a, b, c))
	bytes, _ := sonic.Marshal("hello")
	fmt.Println(string(bytes))
}
