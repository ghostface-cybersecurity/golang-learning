package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 5; i++ {
		func(a int) {
			fmt.Println("Anonymous func: ", a)
		}(i)
	}
}
