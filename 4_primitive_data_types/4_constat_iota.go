package main

import "fmt"

const (
	pi     = 3.1415 //Now this is accessible at the entire package level.
	first  = 1      //CONSTATNS MUST be evaluated at compile time NOT runtime. We cannot set constant = to the result of function.
	second = "second"

	first_iota  = iota // atomatically increments
	second_iota = iota

	//constant expressions

	first_constant_expression  = iota + 6
	second_constant_expression = 2 << iota //bit shift.
	third_constant_expression              //The expression above is re-used here becasue we used iota above.
)

const (
	block_2_iota = iota //iota resets in every constant block. This result is 0.
)

func main() {
	fmt.Println(pi)
	fmt.Println(first, second)

	fmt.Println(first_iota, second_iota)
	fmt.Println(first_constant_expression, second_constant_expression, third_constant_expression)

	fmt.Println(block_2_iota) // iota resets in every block.
}