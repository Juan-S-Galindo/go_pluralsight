// You can edit this code!
// Click here and start typing.
package main

func main() {
	slice := []int{1, 2, 3}

	for indexer, value := range slice { //We define 2 variabels. i: indexer, v: value. value is set to reange keyword and the name of the slice we want to interate.
		println(indexer, value)
	}

	TestMapLoop := map[string]int{"http": 80, "https": 443}

	for k, v := range TestMapLoop {
		println(k, v)
	}

	for k := range TestMapLoop { //If we just want the keys
		println(k)
	}

	for _, v := range TestMapLoop { //If we just want the values.
		println(v)
	}

}
