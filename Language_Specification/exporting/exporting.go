package main

import (
	"fmt"
	"github.com/naraycitra/learn-golang/Language_Specification/exporting/codes"
)

// Exported identifier
func main() {
	// exported identifier
	fmt.Println(codes.ArchiveAlreadyDel)
	// access a value of an unexported identifier
	err100 := codes.New(100)
	fmt.Println(err100)
	// unexported fields from an exported struct
	status := codes.Status{Name: "s"}
	fmt.Println(status)
}
