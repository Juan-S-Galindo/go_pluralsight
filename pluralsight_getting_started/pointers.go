// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {
	// Pointer -> Instead of holding value in var, we hold the address of the location in memory that hold the variable for us

	var firstNameUnInit *string  //* denotes pointer
	fmt.Println(firstNameUnInit) // results in <nil> --> terminology for empty pointer beccause we have not initialized the pointer

	var firstName *string = new(string)
	fmt.Println(firstName) //Prints location in memory

	*firstName = "Sebastian" //In this case we are dereferencing. We convert a pointer to regular var.

	fmt.Println(*firstName) //Prints value stored in the pointer

}

func mai2() {
	//Address of operator

	firstName := "Sebastian"
	fmt.Println(firstName)

	ptr := &firstName
	fmt.Println(ptr, *ptr)

	firstName = "Juan"
	fmt.Println(ptr, *ptr) //location in memory is the same but value stored is different.

}

//deploy layer to nomitesting de dev and DEV
