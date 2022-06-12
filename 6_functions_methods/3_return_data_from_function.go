// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
)

func main() { //main() is special because it defines entry point.
	fmt.Println("Main line 1")
	port := 3000
	num_retries := 3
	port, err := startWebServer(port, num_retries) //We use implicit init to capture the returned value. We can use _ similar to python to ignore variables that are unused without raising the unused error.
	fmt.Println(port, err)                         //Here we can specify what to do with the error if we get an error.

}

func startWebServer(port, numberRetries int) (int, error) { //We can return multiple data types: int or error for example
	fmt.Println("Starting server...")
	fmt.Println("Server Started in port", port)
	fmt.Println("Number of retries", numberRetries)
	return port, nil //If the function does not have to return a value, we can return error. In this way if something goes wrong, we return the error and the calling function can decide what to do with the error.
	//to raise errors use errors.New("Something went wrong")

}
