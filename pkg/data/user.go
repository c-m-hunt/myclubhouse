package data

import (
	"fmt"
)

type User struct {
	ID       int
	Username string
	Title    string
	Forename string
	Surname  string
}

func (u User) FullName() string {
	return fmt.Sprintf("%v %v %v", u.Title, u.Forename, u.Surname)
}
