package statements

import (
	"fmt"
	"os"
)

func for_statements() {
	// for is while
	i := 1
	for i < 3 {
		fmt.Println(i)
		i++
	}

	// traditional for
	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}

	// for range
	for index, value := range []string{"a", "b", "z"} {
		fmt.Println(index, value)
	}

	// for goroutine
	for i := 1; i < 10; i++ {
		fmt.Println("for:", i)
		go func() {
			fmt.Println("go", i)
		}()
	}
}

func if_else_statements() {
	// 3 types if else, this is no ternary in go
	if 4%2 == 0 {
		fmt.Println("this is true if")
	}

	if 4%2 == 1 {
		fmt.Println("this is true if")
	} else {
		fmt.Println("this is false if")
	}

	if 4%2 == 1 {
		fmt.Println("this is true if")
	} else if 4%3 == 1 {
		fmt.Println("this is true if else if")
	} else {
		fmt.Println("this is else if")
	}

	// bad code
	groupBad, errBad := os.Getgroups()
	if errBad == nil {
		fmt.Println(groupBad)
	}

	// good code
	if groups, err := os.Getgroups(); err != nil {
		for _, g := range groups {
			fmt.Println(g)
		}
	}
}

func select_statments() {
	ch1 := make(chan string, 1)
	ch2 := make(chan int, 1)

	ch1 <- "hello"
	ch2 <- 1
	for i := 1; i < 3; i++ {
		// always no keys
		// case send channel or receive channel
		select {
		case msg := <-ch1:
			fmt.Println("this is from ch1", msg)
		case msg := <-ch2:
			fmt.Println("this is from ch2", msg)
		}

		select {
		case <-ch1: // receive expression
		case b := <-ch1:
			{ // identifier list
				fmt.Println(b)
			}
		}
	}
}

func Switch(v interface{}) {
	switch v.(type) {
	case string:
		fmt.Println("this is string")
	case int:
		fmt.Println("this is int")
	case bool:
		fmt.Println("this is bool")
	default:
		fmt.Println("i don't know")
	}
}
