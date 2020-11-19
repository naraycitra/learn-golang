package main

import "fmt"

const (
	A1 = iota // like serial in the database, self-increment, iota default increment by 1
	A2
	A3
)

const (
	B1 = iota + 2 // start by iota + 2, increment by iota
	B2
	B3
)

const (
	C1 = iota * 2 // start by iota * 2, increment by iota
	C2
	C3
	C4
)

const (
	D1 = 1 << iota // start by 1, and then bit shifting by 1 
	D2
	D3
	D4
)

func main() {
	fmt.Println("A1 : %T, [%v] \n", A1, A1)
	fmt.Println(A1, A2, A3)
	fmt.Println(B1, B2, B3)
	fmt.Println(C1, C2, C3, C4)
	fmt.Println(D1, D2, D3, D4)
}
