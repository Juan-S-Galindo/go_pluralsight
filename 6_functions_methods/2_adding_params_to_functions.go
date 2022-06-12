// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {
	fmt.Println("Main line 1")
	port := 3000
	num_retries := 3
	startWebServer(port, num_retries)
	startWebServer2(port, num_retries)

}

func startWebServer(port int, numberRetries int) {
	fmt.Println("Starting server...")
	fmt.Println("Server Started in port", port)
	fmt.Println("Number of retries", numberRetries)

}

func startWebServer2(port, numberRetries int) { //Note that since both arguments are type int, we can just declare int at the end.
	fmt.Println("Starting server...")
	fmt.Println("Server Started in port", port)
	fmt.Println("Number of retries", numberRetries)

}
