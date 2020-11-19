package main

import (
	"fmt"
	"reflect"
)

type Test struct {
	Start string
}

// value receiver
func (t Test) Finish() string {
	return t.Start + "finish"
}

// pointer receiver
func (t *Test) Another() string {
	return t.Start + "another"
}

func CallMethod(i interface{}, methodName string) interface{} {
	var ptr reflect.Value
	var value reflect.Value
	var finalMethod reflect.Value

	value = reflect.ValueOf(i)

	// if we start with a pointer, we need to get value pointed to
	// if we start with a value, we need to get a pointer to that value
	if value.Type().Kind() == reflect.Ptr {
		ptr = value
		value = ptr.Elem()
	} else {
		ptr = reflect.New(reflect.TypeOf(i))
		temp := ptr.Elem()
		temp.Set(value)
	}

	// check for method on value
	method := value.MethodByName(methodName)
	if method.IsValid() {
		finalMethod = method
	}

	// check for method in pointer
	method = ptr.MethodByName(methodName)
	if method.IsValid() {
		finalMethod = method
	}

	if finalMethod.IsValid() {
		return finalMethod.Call([]reflect.Value{})[0].Interface()
	}

	// return or panic, method not found of either type
	return ""
}

func main() {
	i := Test{Start: "start"}
	j := Test{Start: "start2"}

	fmt.Println(CallMethod(i, "Finish"))
	fmt.Println(CallMethod(&i, "Finish"))
	fmt.Println(CallMethod(i, "Another"))
	fmt.Println(CallMethod(&i, "Another"))
	fmt.Println(CallMethod(j, "Finish"))
	fmt.Println(CallMethod(&j, "Finish"))
	fmt.Println(CallMethod(j, "Another"))
	fmt.Println(CallMethod(&j, "Another"))
}
