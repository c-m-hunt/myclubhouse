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
	DOB                            string
	Age                            int
	RegistrationDate               string
	IsCurrentMember                bool
	MembershipCategories           []string
	SubscriptionResources          []string
	GroupName                      string
	MembershipNumber               string
	MembershipExpiryDate           string
	FamilyGroup                    string
	Parent1                        string
	Parent2                        string
	TokenBalance                   int
	LastLoginTime                  string
	Notes                          string
	Occupation                     string
	Attributes                     struct{}
}

// FullName - Returns the full name of the user concatenating the
// title forename and surname
func (u User) FullName() string {
	return fmt.Sprintf("%v %v %v", u.Title, u.Forename, u.Surname)
}
