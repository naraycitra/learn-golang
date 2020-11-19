package main

import (
	"fmt"
	"math"
	"sort"
)

// literals functions
func fooo(a, b int) int {
	return a + b
}

// primary expression

// selectors expression

// method expression

func main() {
	// alternative
	a := 1 | 0
	fmt.Println(a)
	arr := []int{1, 2, 3}
	arr = append(arr, 4)
	for e := range arr {
		fmt.Println(e)
	}

	for k, v := range arr {
		fmt.Println(k, v)
	}

	sort.Ints(arr)
	fmt.Println(arr)
	arrstring := []string{"d", "c", "b", "a"}
	sort.Strings(arrstring)
	fmt.Println(arrstring)
	// package qualified indentifiers
	fmt.Println(math.Max(2, 4))
	// composite literals
	fmt.Println([]int{1, 3})
}
