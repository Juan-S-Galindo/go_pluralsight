// You can edit this code!
// Click here and start typing.
package main

import "fmt"

// MAPS: are collections taht associate keys with values of the SAME DATA TYPE for Key/VAlUE.
func main() {
	m := map[string]int{"foo": 42} //[KEY_data_type]VALUE_data_type
	fmt.Println(m)
	fmt.Println(m["foo"])

	m["foo"] = 28
	fmt.Println(m)
	fmt.Println(m["foo"])

	//How to delete elements

	delete(m, "foo")
	fmt.Println(m)
}