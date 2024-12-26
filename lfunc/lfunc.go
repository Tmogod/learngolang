package main

import (
	"errors"
	"fmt"
)

func foo(a int, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("b is zero")
	}
	d := a / b
	return d, nil
}

func main() {
	m, n := 4, 1

	if q, err := foo(m, n); err != nil {
		fmt.Printf("发生了异常, %v\n", err)

	} else {
		fmt.Printf("商为: %d", q)
	}

}
