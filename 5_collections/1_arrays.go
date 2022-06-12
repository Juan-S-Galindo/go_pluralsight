package main

import "fmt"

//Array: FIXED SIZED OF SIMILAR DATA TYPES. ALl elements MUST have the same type.
func main() {

	//Long syntax: [size]data_type
	var arr [3]int
	arr[0] = 1 //to assign values.
	arr[1] = 2
	arr[2] = 3
	fmt.Println(arr)
	fmt.Println(arr[1])
	fmt.Println(arr[4]) //Compiler will check for bounds to make sure we do not try to get an index outside of the bounds.

	//short form
	arr2 := [3]int{1, 2, 3}
	fmt.Println(arr2)
}