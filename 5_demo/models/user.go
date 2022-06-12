package models

type User struct{
	ID int
	FirstName string
	LastName string
}

var (
	users []*User //Slice of Users which at this point are just pointers.
	nextId = 1 //At package level, we do not need : to do implicit init. By providint the value, the compiler will figure out data_type. 
)