package main

import (
	"fmt"
	"sort"
)

func main() {
	// init array, declare array, initial array, initial two dimensional arrays, multi-line initial array
	// set array, get array, len of array, comparation array, append, sort, iteration array for range
	// declare and initialize
	var a1 [3]int
	fmt.Println("this is a1:", a1)
	// declare two dimensional array
	var a2 [3][3]int
	fmt.Println("this is a2:", a2)
	// declare and initial value, longhand: var a1 [n]T = [n]T{v1,v2,...,vn,}, shorthand: a1 := [n]T{v1,v2,...,vn,}
	a3 := []string{"aa", "bb"}
	fmt.Println("this shorthand array a3:", a3)
	// declare longhand array
	var a4 []string
	fmt.Printf("this is a4: %v, len: %v, cap: %v\n", a4, len(a4), cap(a4))
	a5 := [...]string{"cc", "dd", "ee"}
	fmt.Printf("this is a5: %v, len: %v, cap: %v\n", a5, len(a5), cap(a5))
	a6 := [][]string{{"a1", "a2"}, {"b1", "b2"}}
	fmt.Printf("this is a6: %v, len: %v, cap: %v\n", a6, len(a6), cap(a6))

	// value assignment
	a1[0] = 2
	a2[0][0] = 0
	// append
	a3 = append(a3, "cc")
	//sort
	sort.Strings(a3)
	// copy
	copy(a3, a4)
	// iterate with range
	for i, a := range a3 {
		fmt.Println("iterate with range: index %d, value: %v\n", i, a)
	}
	// iterate with traditional style
	for i := 0; i < len(a3); i++ {
		fmt.Println("iterate with range: index %d, value: %v\n", i, a3[i])
	}

	// contiguous memory allocations
	// declare an array of 6 strings initialized with values
	a7 := [6]string{"Naray", "Nurul", "Abdullah", "Fatimah", "Citra", "Istiqomah"}
	fmt.Printf("\n=> contiguous memory allocation\n")
	for i, v := range a7 {
		fmt.Printf("Value[%s]\tAddress[%p] IndexAddr[%p]\n", v, &v, &a7[i])
	}
}
