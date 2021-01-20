package data_test

import (
	"testing"
	"time"

	"github.com/c-m-hunt/myclubhouse/data"
)

func TestItGetsUsersNotLoggedIn(t *testing.T) {
	oneDayAgo := data.Time(time.Now().AddDate(0, 0, -1))
	threeDaysAgo := data.Time(time.Now().AddDate(0, 0, -3))
	u := data.Users{
		data.User{
			Forename:      "Person1",
			LastLoginTime: &oneDayAgo,
		},
		data.User{
			Forename:      "Person1",
			LastLoginTime: &threeDaysAgo,
		},
		data.User{
			Forename: "Person1",
		},
	}

	li := u.NotLoggedIn(time.Now().AddDate(0, 0, -2))
	if len(li) != 2 {
		t.Fatal("Not correctly getting not logged in people")
	}
}
