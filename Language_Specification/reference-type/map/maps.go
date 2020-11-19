package main

import "fmt"

// user defines a user in program
type user struct {
	name    string
	surname string
}

func main() {
	// declare and initialize

	// declare and make a map that stores values of type user with a key of type string
	users1 := make(map[string]user)

	// add key/value pairs to the map
	users1["naray"] = user{
		name:    "naray",
		surname: "citra",
	}
	users1["nurul"] = user{
		name:    "nurul",
		surname: "istiqomah",
	}
	users1["abdullah"] = user{
		name:    "abdullah",
		surname: "muhammad alfatih",
	}
	users1["fatimah"] = user{
		name:    "fatimah",
		surname: "azzahra khawla",
	}

	// iterate over the map
	fmt.Printf("\n=> Iterate over map with range\n")
	for key, value := range users1 {
		fmt.Println(key, value)
	}

	// map literals

	// declare and initialize the map with values
	users2 := map[string]user{
		"naray":    {"naray", "citra"},
		"nurul":    {"nurul", "istiqomah"},
		"abdullah": {"abdullah", "muhammad alfatih"},
		"fatimah":  {"fatimah", "azzahra khawla"},
	}

	// iterate over the map
	fmt.Printf("\n=>  map literals\n")
	for key, value := range users2 {
		fmt.Println(key, value)
	}

	// delete key
	delete(users2, "naray")

	// find key

	// find the naray key
	// if found is true, we will get a copy value of that type
	// if found is false, we still a value of type user but is to its zero value
	u1, found1 := users2["naray"]
	u2, found2 := users2["nurul"]

	// display the value and found flag
	fmt.Printf("\n=> find key\n")
	fmt.Println("naray", found1, u1)
	fmt.Println("nurul", found2, u2)
}
