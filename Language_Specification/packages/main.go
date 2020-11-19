package main

import (
	"fmt"
	"github.com/naraycitra/learn-golang/Language_Specification/packages/app"
)

/*
 package name a.b.c
 main.go is entry
 package compilation: go build -o main ./main.go
 package install: go install ./main.go
  $GOPATH/bin/main
  package init function: first import, first init
 package init function: init first, then main
 package init function, main job is initialize global variables
  import package: use package alias when import when import package with the some name
  import package: use underscore "_" when only initialize package
  import package: use dot ".", then all export member can call from local file, not very fond of it
 imported package: is initialized only once per package
*/

/*
 Program execution order
 go run *.go
 |- Main package is executed
 |- All imported packages are initialized
 | |- All imported packages are initialized (recursive definition)
 | |- All global variables are initialized
 | -  init functions are called in lexical dile name order
 - Main package is initialized
   |- All global variables are initialized
   -  init functions are called in lexical file name order
*/

func main() {
	fmt.Println("this is main")
	app.App()
}
