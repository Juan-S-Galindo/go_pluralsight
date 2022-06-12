// You can edit this code!
// Click here and start typing.
package main

type User struct {
	ID        int
	FirstName string
	LastName  string
}

func main() {
	u1 := User{
		ID:        1,
		FirstName: "Arthur",
		LastName:  "Morgan",
	}
	u2 := User{
		ID:        1,
		FirstName: "Ford",
		LastName:  "Prefect",
	}
	if u1.ID == u2.ID { //equal
		println("Same user!")
	}
	if u1.FirstName != u2.FirstName { //Not equal
		println("Different name!")
	}

	if u1.ID != u2.ID {
		println("Different IDs!")
	} else { // Else statement
		println("Same IDs")
	}

	if u1 == u2 { //We can compare structs.
		println("same struct!")
	} else if u1.FirstName == u2.FirstName {
		println("Same name")
	} else {
		println("No match for same id and name")
	}

	if u1 == u2 { //We can compare structs.
		println("same struct!")
	} else if u1.FirstName == u2.FirstName {
		println("Same name")
	} //This works as well, it has the option to meet 2 conditions but does not have an else to default. same as python.

}
