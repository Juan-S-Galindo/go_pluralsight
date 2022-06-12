package main

import "fmt"

func main() {
	//Constants cannot be changed after declared

	const pi = 3.1415 //Value needs to be initialized when constant is declared. Value must be resolved at compiled time. Functions are resolved at runtime so we cannot set at constant to a result of a function.
	fmt.Println(pi)

	//pi = 5 -> creates error becase we  cannot change a constant that is already declared.

	fmt.Println(pi + 3)

	const c int = 3
	fmt.Println(float32(c) + 1.2) //We need float32() to convert int to float because we assigned a type at the momment of declaration.

}
