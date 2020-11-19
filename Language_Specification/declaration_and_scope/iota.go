package main

import "fmt"

const (
	c0 = iota // c0 = 1
	c1 = iota // c1 = 2
	c2 = iota // c2 = 3
)

const (
	d0 = iota
	d1
	d2
)

const (
	a = 1 << iota // a == 1 (iota == 0)
	b = 1         // b == 1 (iota == 1, unused)
	c = 3         // c == 3 (iota == 2, unused)
	d = 1         // d == 1 (iota == 3, unused)
)

const (
	aa = 1 << iota // 1
	bb             // 0010 = 2
	cc             // 0100 = 4
	dd             // 1000 = 8
)

const (
	u         = iota * 42 // u == 0 (untyped integer constant)
	v float64 = iota * 42 // v = 42.0 (float64 constant)
	w         = iota * 42 // w = 84 (untyped integer constant)
)

const x = iota // x == 0
const y = iota // y == 0

func main() {
	fmt.Println(c0, c1, c2)
	fmt.Println(d0, d1, d2)
	fmt.Println(a, b, c, d)
	fmt.Println(aa, bb, cc, dd)
	fmt.Println(u, v, w)
	fmt.Println(x, y)
}
