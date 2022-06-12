// You can edit this code!
// Click here and start typing.
package main

import "fmt"

//Slice: Dynamically sized arrays of the same data type.
func main() {
	//We can create slices by using pointers to an existing array
	arr := [3]int{1, 2, 3}

	slice := arr[:]

	arr[1] = 42
	slice[2] = 27 //Because this is a pointer, the underlying array will be modified as well.

	fmt.Println(arr, slice)

	//We can create a slice without underlying array too.

	slice2 := []int{1, 2, 3} //Wihtout size, the slice will be unbounded.Go will manage underlying array.
	fmt.Println(slice2)

	slice2 = append(slice2, 4)
	fmt.Println(slice2)

	s2 := slice2[1:] //We can create slice of slices. Similar to python when we slice lists.
	s3 := slice2[:2]
	s4 := slice2[1:2]
	fmt.Println(s2,s3,s4)
}
