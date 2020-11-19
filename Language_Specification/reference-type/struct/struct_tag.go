package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"time"
)

/*
1. only exported Fields can be json output
2. use tag name when has struct tag name
*/

// struct with no tag
type User struct {
	Name      string
	Password  string
	Phone     int64
	CreatedAt time.Time
	Ignoring  int64
}

// struct with json tag
type User1 struct {
	Name      string    `json:"name"`
	Password  string    `json:"password"`
	Phone     int64     `json:"phone"`
	CreatedAt time.Time `json:"created_at"`
	Ignoring  int64     `json:"-"`
}

// Get Tags
func PrintTags(s interface{}) {
	t := reflect.TypeOf(s).Elem()
	v := reflect.ValueOf(s).Elem()
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		vf := v.Field(i)
		fmt.Printf("Struct PrintTags index: %v, field: %v, value: %v, type: %v, tag: %v\n", i, f.Name, vf.Interface(), vf.Type().Name(), f.Tag.Get("json"))
	}
}

// Depopulate to map
func Depopulate(s interface{}, values map[string]string) {
	st := reflect.TypeOf(s).Elem()
	v := reflect.ValueOf(s).Elem()
	for i := 0; i < st.NumField(); i++ {
		f := st.Field(i)
		key := f.Tag.Get("json")
		values[key] = v.Field(i).String()
	}
}

func main() {
	user := &User{
		Name:      "naray",
		Password:  "password",
		Phone:     123,
		CreatedAt: time.Time{},
	}
	user1 := &User1{
		Name:      "naray",
		Password:  "password",
		Phone:     123,
		CreatedAt: time.Time{},
	}

	// json.Marshal Tags
	out, _ := json.Marshal(user)
	out1, _ := json.Marshal(user1)
	fmt.Println("user marshal:", string(out))
	fmt.Println("user1 marshal:", string(out1))

	// json.Unmarshal Tag
	userJson := `
	{
		"name": "naray json",
		"password": "password",
		"phone": 123
	}
	`
	user2 := &User1{}
	_ = json.Unmarshal([]byte(userJson), user2)
	fmt.Printf("json.Unmarshal Tag: %#v\n", user2)
	// print json tags
	PrintTags(user1)
	userMap := make(map[string]string)
	Depopulate(user1, userMap)
	fmt.Printf("struct populate: %v \n", userMap)
}
