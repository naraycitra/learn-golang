package main

import "fmt"

func main() {
	// range, string, array, slice, map
	r := "string"
	rr := []int{1, 2, 3}
	rr = append(rr, 4)
	rrr := make([]int, 0)
	rrr = append(rrr, 5)
	rrrr := map[string]string{"a": "a"}

	for i, v := range r {
		fmt.Println(i, v)
	}

	for _, v := range r {
		fmt.Println(v)
	}

	for k := range r {
		fmt.Println(k)
	}

	for k, v := range rr {
		fmt.Println(k, v)
	}

	for k, v := range rrr {
		fmt.Println(k, v)
	}

	for k, v := range rrrr {
		fmt.Println(k, v)
	}
}
