package data

import (
	"fmt"
	"time"
)

// Users A slice of user
type Users []User

// Time parsable from input json
type Time struct {
	time.Time
}

// UnmarshalJSON Grab the time from input format to a time object
func (t *Time) UnmarshalJSON(b []byte) error {
	strDate := string(b)
	layouts := map[int]string{
		21: `"2006-01-02T15:04:05"`,
		23: `"2006-01-02T15:04:05.0"`,
		24: `"2006-01-02T15:04:05.00"`,
		25: `"2006-01-02T15:04:05.000"`,
	}
	if strDate == "null" || len(strDate) == 0 {
		return nil
	}
	layout := layouts[len(strDate)]
	tParsed, _ := time.Parse(layout, strDate)
	*t = Time{tParsed}
	return nil
}

// User User data structure
type User struct {
	ID       int
	Username string
	Title    string
	Forename string
	Surname  string
	Photo    string
	Address  struct {
		ID            int
		Street        string
		City          string
		County        string
		Country       string
		Postcode      string
		MapLocation   string
		MapLatitude   float32
		MapLongitude  float32
		MailAddressee string
	}
	HomeTelephone     string
	BusinessTelephone string
	MobileTelephone   string
	Email             string
	Email2            string
	Gender            string
	AbilityLevel      struct {
		ID    int
		Name  string
		Level int
	}
	AbilityLevelAchievedDate string
	SecondAbilityLevel       struct {
		ID    int
		Name  string
		Level int
	}
	SecondAbilityLevelAchievedDate string
	DOB                            *Time
	Age                            int
	RegistrationDate               *Time
	IsCurrentMember                bool
	MembershipCategories           []string
	SubscriptionResources          []string
	GroupName                      string
	MembershipNumber               string
	MembershipExpiryDate           *Time
	FamilyGroup                    string
	Parent1                        string
	Parent2                        string
	TokenBalance                   int
	LastLoginTime                  *Time
	Notes                          string
	Occupation                     string
	Attributes                     struct{}
}

// FullName - Returns the full name of the user concatenating the
// title forename and surname
func (u User) FullName() string {
	return fmt.Sprintf("%v %v %v", u.Title, u.Forename, u.Surname)
}

// NotLoggedIn Returns a slice of users who haven't logged in since
// a date
func (us Users) NotLoggedIn(since time.Time) Users {
	retUsers := Users{}
	for _, u := range us {
		if u.LastLoginTime != nil {
			t := *u.LastLoginTime
			if since.Before(t.Time) {
				retUsers = append(retUsers, u)
			}
		}
	}
	return retUsers
}
