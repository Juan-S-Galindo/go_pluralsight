// You can edit this code!
// Click here and start typing.
package main

func main() {

	//infinite loop method 1:

	var i int
	for { // we can passs for ; ; {} and it works. formatter just deletes this way because for {} is preferred.
		if i > 5 { //Loop will run until this condition is met.
			break
		}
		println(i)
		i++
	}

	var y int
	for { // This way is cleaner. pretty much while loop.
		if i > 5 { //Loop will run until this condition is met.
			break
		}
		println(i)
		y++
	}
}
