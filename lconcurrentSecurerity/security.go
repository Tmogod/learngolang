package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var (
	n    int64
	lock = sync.Mutex{}
)

func foo() {
	for i := 0; i < 100000; i++ {
		//lock.Lock()
		//n++
		//lock.Unlock()
		atomic.AddInt64(&n, 1)
	}
}

func main() {
	group := sync.WaitGroup{}
	group.Add(2)
	go func() {
		defer group.Done()
		foo()
	}()
	go func() {
		defer group.Done()
		foo()
	}()
	group.Wait()
	fmt.Printf("总数%d \n", n)
}
