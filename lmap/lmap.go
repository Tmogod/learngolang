package main

import "fmt"

func main() {
	var m map[string]int
	// m = make(map[string]int, 100) // 预计容量100
	m = map[string]int{"a": 3, "b": 6}
	m["c"] = 9
	m["a"] = 4
	fmt.Printf("%d\n", m["a"])
	delete(m, "a")             // 删除"a"对应的key
	fmt.Printf("%d\n", m["a"]) // key不存在, 打印int默认值0

	if v, exist := m["a"]; exist {
		if exist {
			fmt.Printf("%d\n", v)
		}
	}

	for k, v := range m {
		fmt.Printf("key=%s, value=%d\n", k, v)
	}
}
