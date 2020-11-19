package main

// user represent user in system
type user struct {
	name  string
	email string
}

func main() {
	// pass by value

	// declare variable of type int with a value of 10
	// this value is put on a stack with value of 10
	count := 10

	println("count:\tValue Of[", count, "],\tAddr Of[", &count, "]")

	// Pass the "value of" count
	increment1(count)

	// Printing of the result of count. Nothing has changed.
	println("count:\tValue Of[", count, "],\tAddr Of[", &count, "]")

	// pass the "address of" count
	// this is still considered pass by value, not by reference because the address itself is a value
	increment2(&count)

	// printing out the result of count, count is updated
	println("count:\tValue Of[", count, "],\tAddr Of[", &count, "]")

	// escape analisys

	stayOnStack()
	escapeToHeap()
}

func increment1(incr int) {
	// increment the "value of" incr
	incr++
	println("incr:\tValue Of[", incr, "],\tAddr Of[", &incr, "]")
}

// increment2 declares count as a pointer variable whose value is always an address and points to
// value of type int
// the * here is not an operator. it is part of the type name
// every type that is declared, whether you declare or it is predeclared, you get for free a pointer
func increment2(inc *int) {
	// increment the "value of" count that the "pointer points to"
	// The * is an operator. It tells us the value of the pointer points to
	*inc++
	println("inc:\tValue Of[", inc, "],\tAddr Of[", &inc, "]")
}

// stayOnStack shows how the variable does not escape
func stayOnStack() user {
	// in the stayOnStack stack frame, create a value and initialize it
	u :=
		user{
			name:  "Naray",
			email: "naray@gmail.com",
		}

	// take the value and return it, pass back up to main stack frame
	return u
}

// escapeToHeap shows how the variable escape
func escapeToHeap() *user {
	u := user{
		name:  "Naray",
		email: "naray@gmail.com",
	}
	return &u
}
