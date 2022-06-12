package main

import "fmt"

func main() {
	//Constants cannot be changed after declared

	const pi = 3.1415 //Value needs to be initialized when constant is declared.
	fmt.Println(pi)

	//pi = 5 -> creates error

	fmt.Println(pi + 3)

	const c int = 3
	fmt.Println(float32(c) + 1.2) //We need float32() because we assigned a type at the momment of declaration.

}

//Iota and constant expressions

const (
	pi     = 3.1415 //Now this is accessible at the entire package level.
	first  = 1      //CONSTATNS MUST be evaluated at compile time NOT runtime. We cannot set constant = to the result of function.
	second = "second"

	first_iota  = iota // atomatically increments
	second_iota = iota

	//constant expressions

	first_constant_expression  = iota + 6
	second_constant_expression = 2 << iota
	third_constant_expression  //The expression above is used here
)

const (
	block_2_iota = iota
)

func main2() {
	fmt.Println(pi)
	fmt.Println(first, second)

	fmt.Println(first_iota, second_iota)
	fmt.Println(first_constant_expression, second_constant_expression, third_constant_expression)

	fmt.Println(block_2_iota) // iota resets in every block.
}
