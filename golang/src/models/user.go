package models

// User - basic representation of a user
type User struct {
	ID        int
	FirstName string
	LastName  string
}

// you can also have var blocks (like import and constant blocks)
var (
	// this is going to be a slice that holds pointers to user objects
	// the reason for pointers is that it allows us to not have to copy user objects throughout and mainpulate the same user in memory
	users []*User

	// at the PACKAGE level, don't need := syntax for implicit initialization
	nextID = 1
)
