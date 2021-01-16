package cli

import (
	"fmt"
	"log"
	"sort"
	"time"

	"github.com/c-m-hunt/myclubhouse/apiclient"
	"github.com/c-m-hunt/myclubhouse/data"
)

var dateLayout string = "Jan 2nd 2006"

// DisplayExpiringMembers Displays users which are due to expire before specified date
func DisplayExpiringMembers(cfg *Config, before time.Time) {
	printTitle(fmt.Sprintf("EXPIRING MEMBERS by %v", before.Format(dateLayout)))
	users := getUsers(cfg)
	sort.Slice(users, func(i, j int) bool {
		if users[i].MembershipExpiryDate == nil {
			return true
		}
		if users[j].MembershipExpiryDate == nil {
			return false
		}
		return users[i].MembershipExpiryDate.Before(users[j].MembershipExpiryDate.Time)
	})
	for _, u := range users {
		if u.MembershipExpiryDate != nil && u.MembershipExpiryDate.Before(before) {
			fmt.Printf("%7v %-35v %20v\n", u.ID, u.FullName(), u.MembershipExpiryDate.Format(dateLayout))
		}
	}
}

func getUsers(cfg *Config) data.Users {
	c := apiclient.MakeClient(cfg.SubDomain, cfg.AccessToken, nil)

	q := apiclient.RequestQuery{
		PageSize: 200,
	}

	users := []data.User{}
	_, err := c.Users(&q, &users)

	if err != nil {
		log.Fatal(err)
	}
	return users
}
