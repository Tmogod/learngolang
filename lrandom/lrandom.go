package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	//source := rand.NewPCG(123, 456)
	//for i := 0; i < 5; i++ {
	//	rander := rand.New(source)
	//	fmt.Printf("%d ", rander.IntN(100))
	//}

	fmt.Printf("%d  ", rand.IntN(100))
}
