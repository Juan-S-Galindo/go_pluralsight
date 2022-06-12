// You can edit this code!
// Click here and start typing.
package main

//Panic: is a situation similar to Exception in python. It happens when it has no idea how to proceed.
func main() {
	println("Starting web server")

	panic("Something just bad happened") // To learn how to recover from panic, read golang.org/doc/effective_go.html#panic

	println("Web server started")
}
