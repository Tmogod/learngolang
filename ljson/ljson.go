package main

import (
	"encoding/json"
	"fmt"
	"github.com/bytedance/sonic"
)

type Student struct {
	Name   string
	Age    int
	Gender bool
}

type Class struct {
	Id       string
	Students []Student
}

func main() {
	f := Student{Name: "张三", Age: 18, Gender: true}
	c := Class{
		Id:       "1(2)",
		Students: []Student{f, f, f},
	}
	marshal, err := json.Marshal(c)
	if err != nil {
		return
	}
	fmt.Println(string(marshal))

	d := Class{}

	err = json.Unmarshal(marshal, &d)
	if err != nil {
		fmt.Println("json反序列化失败", err)
		return
	}
	fmt.Printf("%v\n", d)

	marshal1, err := sonic.Marshal(c)
	if err != nil {
		return
	}

	err = sonic.Unmarshal(marshal1, &d)
	if err != nil {
		fmt.Println("json反序列化失败", err)
		return
	}
}
