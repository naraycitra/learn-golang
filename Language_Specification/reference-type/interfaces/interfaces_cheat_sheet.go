package main

import (
	"fmt"
	_ "net/http/pprof"
)

/*
 interface: the simplest interface is empty interface{}, all type are inherited to it
 interface: error is also an intefrece, including Error method
 interface: is a collection of method, and only method
 interface: can be implement
 interface: only provide method name, argument, return type, method defined is declare in type
 interface: type is nil, value is nil when defined, only have type and value when implement
 interface: the value and type is dynamic, is the type and value of implement type
 interface: a struct can implement multi interface
*/

// declaring interface type
type RectInterface interface {
	area() float64
	abc() float64
}

type NameInterface interface {
	getName() string
}

// declaring struct type
type Rect struct {
	width  float64
	height float64
	name   string
}

type Circle struct {
	radius float64
}

// declaring Rect struct method
func (r *Rect) area() float64 {
	return r.width * r.height
}

func (r *Rect) abc() float64 {
	return r.width * r.height
}

func (r *Rect) getName() string {
	return r.name
}

// declaring Circle struct method
func (c *Circle) area() float64 {
	return c.radius * c.radius
}

func (c *Circle) abc() float64 {
	return c.radius * c.radius
}

// Polimorphic function
func GetArea(r RectInterface) float64 {
	fmt.Printf("Polymorphic function, Typ:L %T, area: %v\n", r, r.area())
	return r.area()
}

func main() {
	r := &Rect{width: 3, height: 3, name: "bujur sangkar"}
	fmt.Printf("this is Rect, Type: %T, area: %v, abc:%v \n, getName: %v\n", r, r.area(), r.abc(), r.getName())
	GetArea(r)

	var ri RectInterface
	ri = &Rect{width: 3, height: 3}
	GetArea(ri)
	fmt.Printf("this is RectInterface, Rect Struct, Type: %T, area: %v, abc: %v, nogetname\n", ri, ri.area(), ri.abc())

	ri = &Circle{radius: 3}
	GetArea(ri)
	fmt.Printf("this is RectInterface, Circle Struct, Type: %T, area: %v, abc:%v, nofetname\n", ri, ri.area(), ri.abc())

	ni := &Rect{width: 3, height: 3, name: "segi empat"}
	fmt.Printf("this is NameInterface, Rect Struct, Type: %T, noarea, noabc, getName: %v\n", ni, ni, ni.getName())

	// interface Type assertions
	var i interface{} = "hello"
	fmt.Printf("i am string: %T, %v\n", i, i)
	is, ok := i.(string)
	fmt.Printf("i am string after type assertions: %T, %v, %v\n", is, is, ok)
	i = 1
	it, _ := i.(int)
	fmt.Printf("i am int after type assertions: %T, %v\n", it, it)
}
