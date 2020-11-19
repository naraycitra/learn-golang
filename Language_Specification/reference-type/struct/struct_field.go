package main

import (
	"fmt"
	"reflect"
)

/*
struct: field
struct: field function
struct: method: In this way, we can add method with the same name because the receive is different
struct: type definition and method definition need to be in the same package
*/

type Range struct {
	From int64
	To   int64
}

type Owner struct {
	Id    int64
	Phone int64
}

type Company struct {
	A int32
	B int64
	C string
	Owner
	TagIds []string
	Price  Range
	F1     func(s string) string // field function generally does not do this, it does not conform to the concept of OOP

}

func (s *Company) PrintCompany() { // struct method
	fmt.Println("this is receive method1", s)
}

func main() {
	// Initialize
	company := &Company{
		A: int32(32),
		B: int64(64),
		C: "skadoosh",
		Owner: Owner{
			Id:    int64(6464),
			Phone: int64(8888),
		},
		TagIds: []string{"a1", "a2"},
		Price: Range{
			From: int64(2),
			To:   int64(4),
		},
		F1: func(s string) string {
			return s
		},
	}

	// Selector and promoted
	fmt.Println(company.A, company.Id, company.Phone)
	// iterate struct field name and value
	v := reflect.ValueOf(company).Elem()
	for i := 0; i < v.NumField(); i++ {
		f := v.Field(i)
		fmt.Printf("index : %d, name : %s, type : %s, value : %v\n", i, f.Type().Name(), f.Type(), f.Interface())
	}
	// call function
	fmt.Println(company.F1("m1"))
	// call method
	company.PrintCompany()
}
