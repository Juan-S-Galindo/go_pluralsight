// You can edit this code!
// Click here and start typing.
package main

func main() {

	var i int   //If we declare here, we can use i outside the scope of the loop.
	for i < 5 { //Loop will run until the condition is met.
		println(i)
		i++

		if i == 3 {
			break //We can use break to finish loop if we meet another condition.
		}

		if i == 1 {
			println("continue...")
			continue //We can use break to finish loop if we meet another condition.
		}

	}
}
