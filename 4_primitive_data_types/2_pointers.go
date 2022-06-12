package main

import "fmt"

func main() {
	// Pointer -> Instead of holding value in var, we hold the address of the location in memory that hold the variable for us

	var firstNameUnInit *string  //* denotes pointer, in this case we have a pointer to a location in memory that will store a string value.
	fmt.Println(firstNameUnInit) // results in <nil> --> terminology for empty pointer because we have not initialized the pointer

	var firstName *string = new(string) //This is how we initialize a pointer.
	*firstName = "Sebastian" //to dereference we use *. We use dereference to assign and view values stored in a pointer.

	fmt.Println(firstName) //Prints address in memory
	fmt.Println(*firstName) //Prints value stored in the pointer

}

func mai2() {
	//Address of operator

	firstName := "Sebastian"
	fmt.Println(firstName)

	ptr := &firstName //Address of operator == &. This creates a pointer.
	fmt.Println(ptr, *ptr)

	firstName = "Juan"
	fmt.Println(ptr, *ptr) //location in memory is the same but value stored is different.

}
