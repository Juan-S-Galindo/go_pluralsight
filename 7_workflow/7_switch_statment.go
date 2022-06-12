// You can edit this code!
// Click here and start typing.
package main

type HTTPRequest struct {
	Method string
}

func main() {
	r := HTTPRequest{Method: "GET"}

	switch r.Method {
	case "GET":
		println("Get request") //Note we do not have to use break
		fallthrough            //Executes case below.
	case "DELETE":
		println("DELETE request")
	case "POST":
		println("POST request")
	case "PUT":
		println("PUT request")
	default:
		println("default") //If none of the cases executes, then we can use default.
	}

}
