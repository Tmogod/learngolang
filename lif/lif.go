package main

import "fmt"

func main() {
	// a,b := 10, true
	//a := 10
	//b := true
	if a, b := 10, true; a > 5 {
		fmt.Println("a > 5")
	} else if b {
		fmt.Println("b => true")
	} else {
		fmt.Println("final")
	}
}
