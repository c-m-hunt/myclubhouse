package apiclient

import (
	"fmt"
	"log"

	"github.com/c-m-hunt/myclubhouse/data"
	"github.com/google/go-querystring/query"
)

// UsersResponse - Response back from the client on requesting users
type UsersResponse struct {
	Users      []data.User
	Pagination ResponsePagination
}

// EventsResponse - Response back from the client on requesting events
type EventsResponse struct {
	Events     []data.Event
	Pagination ResponsePagination
}

// Users - Get users from the api client
func (c Client) Users(rq *RequestQuery) (*[]data.User, error) {
	ur, err := c.RequestUsers(rq)
	if err != nil {
		return nil, err
	}
	return &ur.Users, nil
}

// User - Retrieves an individual user based on ID
func (c Client) User(id int) (*data.User, error) {
	url := fmt.Sprintf("%v%v/%v", c.BaseURL, "users", fmt.Sprint(id))
	userResponse := data.User{}
	err := c.getResponse(url, &userResponse)
	if err != nil {
		return nil, err
	}
	return &userResponse, nil
}

// Events - Get events from the api client
func (c Client) Events(rq *RequestQuery) (*[]data.Event, error) {
	er, err := c.RequestEvents(rq)
	if err != nil {
		log.Fatal(err)
	}
	return &er.Events, err
}

// Event - Retrieves an individual event
func (c Client) Event(id int) (*data.Event, error) {
	url := fmt.Sprintf("%v%v/%v", c.BaseURL, "events", fmt.Sprint(id))
	eventResponse := data.Event{}
	err := c.getResponse(url, &eventResponse)
	if err != nil {
		return nil, err
	}
	return &eventResponse, nil
}

// RequestUsers - Requests users from the client
func (c Client) RequestUsers(rq *RequestQuery) (*UsersResponse, error) {
	v, _ := query.Values(rq)
	url := fmt.Sprintf("%v%v?%v", c.BaseURL, "users", v.Encode())
	usersResponse := UsersResponse{}
	err := c.getResponse(url, &usersResponse)
	if err != nil {
		return nil, err
	}
	return &usersResponse, nil
}

// RequestEvents - Requests events from the client
func (c Client) RequestEvents(rq *RequestQuery) (*EventsResponse, error) {
	v, _ := query.Values(rq)
	url := fmt.Sprintf("%v%v?%v", c.BaseURL, "events", v.Encode())
	eventsResponse := EventsResponse{}
	err := c.getResponse(url, &eventsResponse)
	if err != nil {
		return nil, err
	}
	return &eventsResponse, nil
}
