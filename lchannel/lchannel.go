package main

import (
	"fmt"
	"sync"
	"time"
)

func foo() {
	ch := make(chan int, 1) // 容量为0的话，写入即阻塞
	fmt.Println(time.Now().Unix())
	go func() {
		a := <-ch
		fmt.Printf("读出%d, %d\n", a, time.Now().Unix())
	}()

	time.Sleep(1 * time.Second)

	go func() {
		ch <- 4
		fmt.Println(time.Now().Unix(), "写入4成功")
	}()

	time.Sleep(1 * time.Second)
}

func main() {
	ch := make(chan int, 100)

	// 2个生产者，往channel里写入元素
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()

	// 一个消费者
	mc := make(chan struct{}, 0) // 不关心元素，可以使用空结构体来替代

	go func() {
		sum := 0
		for {
			a, ok := <-ch // channel关闭 且为空，此时ok才为false
			if !ok {
				fmt.Println(sum)
				break
			} else {
				sum += a
			}
		}
		mc <- struct{}{}
	}()

	wg.Wait()
	close(ch)
	<-mc
}
