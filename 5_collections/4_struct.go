// You can edit this code!
// Click here and start typing.
package main

import "fmt"

//STRUCT: Only collection that allow us to store disparate data types. FIELDS are fixed at compiled time.
func main() {

	//long way.
	type user struct { //Definition of user. This declaration can be done at the package level.
		ID        int
		FirstName string
		LastName  string
	}

	var u user //Initialization of user struct
	u.ID = 1
	u.FirstName = "Sebastian"
	u.LastName = "Galindo"
	fmt.Println(u)

	//short way

	u2 := user{ID: 1,
		FirstName: "Sebastian",
		LastName:  "Galindo",
	}
	fmt.Println(u2)
}
