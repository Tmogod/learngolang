package main

import (
	"fmt"
	"time"
)

func f2() {
	time.Sleep(1 * time.Second)
	fmt.Println("f2 finish")
}

func f1() {
	fmt.Println("f1 finish")
	go f2()
}

func f3() {
	defer func() {
		err := recover() // 隔离panic风险, 需要放在匿名函数中才生效
		if err != nil {
			fmt.Println(err)
		}
	}()
	a, b := 3, 0
	fmt.Println(a, b)
	_ = a / b // 此处会发生panic(), 发生panic调用defer
	fmt.Println("f1 finish")
}

func main() {
	//wg := sync.WaitGroup{}
	//wg.Add(2)
	//go func() {
	//	defer wg.Done()
	//	fmt.Println("go routine 1")
	//}()
	//go func() {
	//	defer wg.Done()
	//	fmt.Println("foo anonymous func")
	//}()
	//wg.Wait()

	//f1()     // f1执行完，不影响f2的执行
	//time.Sleep(3 * time.Second)

	go f3()
	time.Sleep(1 * time.Second)
	fmt.Println("main finish")
}
