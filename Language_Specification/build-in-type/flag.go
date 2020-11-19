package main

import (
	"flag"
	"fmt"
)

var (
	wordP string
	numP  int
)

func main() {
	flag.StringVar(&wordP, "word", "", "input a string")
	flag.IntVar(&numP, "num", 0, "input a number")
	flag.Parse()
	fmt.Println("word:", wordP)
	fmt.Println("number:", numP)

	a := []string{}
	fmt.Println(len(a))
}
