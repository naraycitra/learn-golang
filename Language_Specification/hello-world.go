package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"time"
)

func Del(keys ...interface{}) {
	fmt.Println(keys...)
}

func main() {
	fmt.Println("hello world")
	fmt.Println(time.Now())
	fmt.Println("this is rand int", rand.Intn(10))
	fmt.Println(math.Sqrt(8))
	fmt.Println(math.Pi)
	pwd, _ := os.Getwd()

	fmt.Println(pwd)
	Del("a", 1, 3)

	a := new(struct {
		B error
	})
	fmt.Println(a)
	if a == nil || a.B == nil {
		fmt.Println("err")
	}
	fmt.Println(2 / 55)
	fmt.Println(fmt.Sprintf("success: %0.2f%%", 3/float32(7)))
}
