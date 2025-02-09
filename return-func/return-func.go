package main

import (
	"fmt"
)

func makeMult(base int) func(int) int {
	return func(factor int) int {
		return base * factor
	}
}

func main() {
	base1 := makeMult(2)
	base2 := makeMult(5)
	for i := 0; i < 3; i++ {
		fmt.Println(base1(i), base2(i))
	}
}
