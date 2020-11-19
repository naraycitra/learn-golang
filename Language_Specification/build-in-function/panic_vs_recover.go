package main

import (
	"fmt"
	"github.com/naraycitra/learn-golang/Language_Specification/errors"
)

// panic, stop execution the current function, then unwinding the stack, then run deferred functions
// panic,
// recover, control of the goroutine and execution
// recover, defere function, unwinding the stact defer function code

func panic1() {
	panic("error in panic1")
}

func panic2() {
	panic("error in panic2")
}

func main() {
	defer func() {
		if errRecover := recover(); errRecover != nil {
			var err error
			isError, ok := errRecover.(error)
			if ok {
				err = isError
			} else {
				err = errors.New(1)
			}
			fmt.Println(err)
		}
	}()
	panic("error in main")
	recover() // unreachable code
	panic1()  // unreachable code
	panic2()  // unreachable code
}
