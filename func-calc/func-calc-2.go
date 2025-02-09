package main

import (
	"fmt"
	"math"
	"strconv"
)

func add(a, b int) int     { return a + b }
func sub(a, b int) int     { return a - b }
func mul(a, b int) int     { return a * b }
func div(a, b int) int     { return a / b }
func div_rem(a, b int) int { return a % b }
func pow(a, b int) int     { return int(math.Pow(float64(a), float64(b))) }

type opFuncType func(int, int) int

var OpMap = map[string]opFuncType{
	"+": add,
	"-": sub,
	"*": mul,
	"/": div,
	"%": sub,
	"^": pow,
}

func main() {
	expressions := [][]string{
		[]string{"2", "+", "3"},
		[]string{"2", "^", "3"},
		[]string{"2", "*", "3"},
		[]string{"2", "-", "2"},
		[]string{"21", "%", "5"},
		[]string{"2", "^", "8"},
		[]string{"50", "%", "2"},
		[]string{"2", "+", "2"},
		[]string{"2", "-", "3"},
		[]string{"70", "*", "3"},
		[]string{"2", "^", "20"},
		[]string{"2", "+", "9"},
		[]string{"50", "^", "2"},
	}
	for _, exp := range expressions {
		if len(exp) != 3 {
			fmt.Println("Error: ", exp)
			continue
		}
		p1, err := strconv.Atoi(exp[0])
		if err != nil {
			fmt.Println(err)
			continue
		}

		op := exp[1]
		opFunc, ok := OpMap[op]
		if !ok {
			fmt.Println("Error: ", op)
		}

		p2, err := strconv.Atoi(exp[2])
		if err != nil {
			fmt.Println(err)
			continue
		}

		result := opFunc(p1, p2)
		fmt.Println(result)
	}
}
