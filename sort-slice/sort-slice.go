package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
	Info string
}

func main() {
	peoples := []Person{
		{"Bob", 47, "engineer"},
		{"Tyler", 22, "software"},
		{"Amber", 32, "developer"},
	}
	fmt.Println(peoples)

	sort.Slice(peoples, func(i int, j int) bool { return peoples[i].Name < peoples[j].Name })
	fmt.Println(peoples)

	sort.Slice(peoples, func(i int, j int) bool { return peoples[i].Age < peoples[j].Age })
	fmt.Println(peoples)
}
