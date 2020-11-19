package main

import "fmt"

func Foo() {
	defer finished()
	defer func() {
		fmt.Println("defer 1")
	}()
	defer func() {
		fmt.Println("defer 2")
	}()
	fmt.Println("start")
	if true {
		return
	}
	nums := []int{1, 3, 4, 5}
	for i, v := range nums {
		fmt.Println(i, v)
	}
}

func finished() {
	fmt.Println("defer finished")
}

func main() {
	Foo()
}
