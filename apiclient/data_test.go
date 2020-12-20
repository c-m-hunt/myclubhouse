// +build integration

package apiclient_test

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/c-m-hunt/myclubhouse/apiclient"
	"github.com/c-m-hunt/myclubhouse/data"
	"github.com/joho/godotenv"
)

var accessToken string
var subdomain string
var c apiclient.Client

var workingUserID int
var workingEventID int

func init() {
	err := godotenv.Load("./../.env")
	if err != nil {
		log.Fatal("Could not load .env")
	}
	accessToken = os.Getenv("ACCESS_TOKEN")
	subdomain = os.Getenv("SUBDOMAIN")
	c = apiclient.MakeClient(subdomain, accessToken, nil)
}

func TestItGetsBasicUsers(t *testing.T) {
	q := apiclient.RequestQuery{
		PageSize: 1,
	}
	users := []data.User{}
	_, err := c.Users(&q, &users)
	if err != nil {
		t.Errorf("Getting users produced an error %v", err)
	}
	if len(users) != 1 {
		t.Errorf("Expected 1 user returned. Got %v", len(users))
	}
	workingUserID = (users)[0].ID
}

func TestItGetsAUser(t *testing.T) {
	user, err := c.User(workingUserID)
	if err != nil {
		t.Errorf("Getting user with ID %v produced an error %v", workingUserID, err)
	}
	fmt.Print(user)
}

func TestItGetsBasicEvents(t *testing.T) {
	q := apiclient.RequestQuery{
		PageSize: 1,
	}
	events := []data.Event{}
	_, err := c.Events(&q, &events)
	if err != nil {
		t.Errorf("Getting events produced an error %v", err)
	}
	if len(events) != 1 {
		t.Errorf("Expected 1 event returned. Got %v", len(events))
	}
	workingEventID = (events)[0].ID
}

func TestItGetsAnEvent(t *testing.T) {
	event, err := c.Event(workingEventID)
	if err != nil {
		t.Errorf("Getting event with ID %v produced an error %v", workingEventID, err)
	}
	fmt.Print(*event)
}
