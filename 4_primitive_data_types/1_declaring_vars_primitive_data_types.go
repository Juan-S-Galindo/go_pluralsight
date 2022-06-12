// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {
	var i int //We can initialize var and assign value later.
	i = 42
	fmt.Println(i)

	var f float32 = 3.14 // We initialize and assign value. We have to specify the size (float32 and float64)
	fmt.Println(f)

	firstName := "Sebastian" //Allows implicit initilization. Go will imply the data type.
	fmt.Println(firstName)

	//var not_used int --> If this is uncomented,it will create error because is not used for anything.

	boolean_type := true
	fmt.Println(boolean_type)

	complex_data_type := complex(3, 4)
	fmt.Println(complex_data_type) // complex numbers are built in

	r, im := real(complex_data_type), imag(complex_data_type)
	fmt.Println(r, im) // can break up real and imaginary components

	i = 1000
	fmt.Println(i)

}
