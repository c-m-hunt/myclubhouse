package apiclient_test

import (
	"log"
	"os"
	"testing"

	"github.com/c-m-hunt/myclubhouse/apiclient"
	"github.com/joho/godotenv"
)

var accessToken string
var subdomain string
var c apiclient.Client

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
	users, err := c.Users(&q)
	if err != nil {
		t.Errorf("Getting users produced an error %v", err)
	}
	if len(*users) != 1 {
		t.Errorf("Expected 1 user returned. Got %v", len(*users))
	}
}

func TestItGetsBasicEvents(t *testing.T) {
	q := apiclient.RequestQuery{
		PageSize: 1,
	}
	events, err := c.Events(&q)
	if err != nil {
		t.Errorf("Getting events produced an error %v", err)
	}
	if len(*events) != 1 {
		t.Errorf("Expected 1 event returned. Got %v", len(*events))
	}
}
