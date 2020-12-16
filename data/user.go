package data

import (
	"fmt"
)

// User - Basic structure of a user
type User struct {
	//TODO(c-m-hunt): Add more fields as required
	ID       int
	Username string
	Title    string
	Forename string
	Surname  string
}

// FullName - Returns the full name of the user concatenating the
// title, forename and surname
func (u User) FullName() string {
	return fmt.Sprintf("%v %v %v", u.Title, u.Forename, u.Surname)
}
