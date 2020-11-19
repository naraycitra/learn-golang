package main

import (
	"fmt"
	"reflect"
)

// ToMap convert a struct to a map using struct's tag
// ToMap uses tags on struct fields to decide which fields to add
// the returned map
func ToMap(in interface{}) (map[string]interface{}, error) {
	out := make(map[string]interface{})

	v := reflect.ValueOf(in)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	// we can only accept struct
	if v.Kind() != reflect.Struct {
		return nil, fmt.Errorf("ToMap can only accept struct; got %T", v)
	}

	typ := v.Type()
	for i := 0; i < v.NumField(); i++ {
		fi := typ.Field(i)
		if tagv := fi.Name; tagv != "" {
			out[tagv] = v.Field(i).Interface()
		}
	}
	return out, nil
}

type Foo struct {
	A int    `gorm:"column:a" json:"a"`
	B string `gorm:"column:b" json:"b"`
	C string
}

func main() {
	f := Foo{
		A: 1,
		B: "foo",
		C: "boo",
	}
	fmt.Printf("%[1]T, %+[1]v\n", f)

	g, err := ToMap(f)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%[1]T, %+[1]v\n", g)
}
