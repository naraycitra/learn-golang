package main

import "fmt"

type Address struct {
	name    string
	street  string
	city    string
	state   string
	zipcode int
}

func main() {
	type user struct{}
	type user1 user   // new type user1
	type user2 = user // user2 is alias user

	// defining name type
	type shop struct {
		Id      string `json:"id"`
		Name    string `json:"name"`
		Owner   string `json:"owner"`
		Orders  []string
		Address Address `json:"address"`
	}
	// defining anonymous type
	shopAnonymous := struct {
		Id    string `json:"id"`
		Name  string `json:"name"`
		Owner string `json:"owner"`
	}{
		Id:    "1",
		Name:  "Anonymous",
		Owner: "Owner",
	}
	fmt.Printf("shopAnonymous: %T, [%v]\n", shopAnonymous, shopAnonymous)

	// initializing a struct with zero value
	shop0 := shop{}
	fmt.Println("Initializing struct with zero value", shop0)

	// initializing a struct without field name
	shop1 := shop{"1", "noname", "owner", []string{}, Address{}}
	fmt.Println("Initializing a struct without field name", shop1)

	// initializing a struct with field name
	shop2 := shop{
		Id:    "2",
		Name:  "name",
		Owner: "owner",
	}
	fmt.Println("Initializing a struct with field name", shop2)

	// new a struct type with zero value
	shop3 := new(shop)
	fmt.Println("new struct type with zero value", shop3)
}
