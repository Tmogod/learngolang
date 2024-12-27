package main

import (
	"encoding/json"
	"fmt"
	"github.com/bytedance/sonic"
	"testing"
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

var (
	s = Student{"张三", 18, true}
	c = Class{
		Id:       "1(2)班",
		Students: []Student{s},
	}
)

func TestJson(t *testing.T) {

	marshal, err := json.Marshal(c)
	if err != nil {
		t.Fail()
	}
	var c2 Class
	err = json.Unmarshal(marshal, &c2)
	if err != nil {
		t.Fail()
	}
	if !(c.Id == c2.Id && len(c.Students) == len(c2.Students)) {
		t.Fail()
	}
	fmt.Printf("json no problem")
}

func TestSonic(t *testing.T) {
	marshal, err := sonic.Marshal(c)
	if err != nil {
		t.Fail()
	}
	var c2 Class
	err = sonic.Unmarshal(marshal, &c2)
	if err != nil {
		t.Fail()
	}
	if !(c.Id == c2.Id && len(c.Students) == len(c2.Students)) {
		t.Fail()
	}
	fmt.Printf("sonic no problem")
}

func BenchmarkJson(b *testing.B) {
	for i := 0; i < b.N; i++ {
		marshal, _ := json.Marshal(c)

		var c2 Class
		_ = json.Unmarshal(marshal, &c2)
	}
}

func BenchmarkSonic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		marshal, _ := sonic.Marshal(c)

		var c2 Class
		_ = sonic.Unmarshal(marshal, &c2)
	}
}
