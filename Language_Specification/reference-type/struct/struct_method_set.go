package main

import "fmt"

// data in struct to bind methods to .
type data struct {
	name string
	age  int
}

// displayName provides a pretty print view of the name.
// it uses data as value receiver.
func (d data) displayName() {
	fmt.Printf("My Name is %s\n", d.name)
}

// setAge sets the age and displays the value.
// it uses data as a pointer receiver.
func (d *data) setAge(age int) {
	d.age = age
	fmt.Println(d.name, "age is", d.age)
}

func main() {
	// Method are just functions

	// declare a variable of type data
	d := data{
		name: "Naray",
	}

	fmt.Println("Proper call methods:")

	// How we actually call methods in Go.
	d.displayName()
	d.setAge(39)

	fmt.Println("\nWhat the Compiler is Doing:")

	// this is what Go is doing underneath
	// when we call d.displayName(), the compiler will call data.displayName, showing that we are
	// using value receiver of type data, and pass the data in as the first parameter,
	// taking a look at the function again: "func (d data) displayName()", that receiver is the
	// parameter because it is trully parameter. it is the first parameter to a function that call
	// displayName.
	// similar to d.setAge(39). Go is calling a function that based on the pointer receiver
	// passing data to its parameters. We are adjusting to make the call by taking the address of d.
	data.displayName(d)
	(*data).setAge(&d, 39)

	// Function variable
	fmt.Println("\nCall Value Receiver Methods with Variable:")

	// Declare a function variable for the method bound to the d variable.
	// The Function variable will get its own copy of d because the method is using a value receiver.
	// f1 is now a reference type: a pointer variable.
	// We don't call the method here. There is no () at the end of displayName.
	f1 := d.displayName

	// call the method via the variable.
	// f1 is a pointer and it points to the special 2 word data structures. This first word points to
	// the code for that method we want to execute, which is displayName in this case. We cannot call
	// displayName unless we have a value of type data. So the second word is a pointer to the copy of
	// data. displayName uses a value reveicer so it works on its own copy. When we make an assignment
	// to f1, we are having a copy of d.
	// | * | --> code
	// | * | --> copy of d
	f1()

	// When we change the value of d to "Naray Citra", f1 is not going to see the change.
	d.name = "Naray Citra"

	// call the method via the variable. We don't see the change.
	f1()

	// however, if we do this again if f2, then we see the change.
	fmt.Println("\nCall Pointer Receiver Method via Variable:")

	// declare a function variable for the method bound to the d variable.
	// the function variable will get the address of d because the method is using a pointer receiver.
	f2 := d.setAge

	// call the method via the variable
	// f2 is also a pointer that has 2 word data structures. The first word points to setAge, but
	// the second words doesn't point to its copy anymore, but its original.
	// | * | --> code
	// | * | --> original d

	// change the value of d
	d.name = "Naray Citra"

	// call method via the variable. We can see the changes
	f2(39)
}
