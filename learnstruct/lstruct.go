package main

import "fmt"

type Human struct {
	name   string
	Age    int
	Height float64
	Sex    bool
}

func main() {
	//var a Human
	//a = Human{name: "fupeng", Age: 12, Height: 150.12, Sex: true}
	a := Human{name: "gaoshanwen", Age: 13, Height: 160.12, Sex: true}
	fmt.Printf("%s %d岁 %f高 %t是男人\n", a.name, a.Age, a.Height, a.Sex)
	fmt.Printf("%v\n", a)  // 详细
	fmt.Printf("%+v\n", a) // 更详细
	fmt.Printf("%#v\n", a) // 更更详细
	//a.name = "lock"
	//a.Age = 12
	//a.Height = 150.12
	//a.Sex = true
}
