package main //package main, it has a unique meaning in Go because package main specifies an entry point for your application

/*
So a module in Go is this entire project space, so all of the source code that's associated with this project space, so everything from this go.mod file,
everything from the directory that contains it on down, that's our module. Within a module, we have what are called packages.
And packages are just discrete units of code that are associated to some sort of a concept within our module.
So in this case, I want to create a package within our module, and we're going to call that models. So this is going to be the model layer for our web service
*/

import (
	"fmt"
	"github.com/pluralsight/webservice"
)

func main() {
	u := models.User{
		ID:2,
		FirstName: "Sebastian",
		LastName: "Galindo",
	}
	fmt.Println(u)
}
