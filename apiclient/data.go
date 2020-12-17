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
	ur, err := c.requestUsers(rq)
	if err != nil {
		return nil, err
	}
	return &ur.Users, nil
}

// Events - Get events from the api client
func (c Client) Events(rq *RequestQuery) (*[]data.Event, error) {
	er, err := c.requestEvents(rq)
	if err != nil {
		log.Fatal(err)
	}
	return &er.Events, err
}

// RequestUsers - Requests users from the client
func (c Client) requestUsers(rq *RequestQuery) (*UsersResponse, error) {
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
func (c Client) requestEvents(rq *RequestQuery) (*EventsResponse, error) {
	v, _ := query.Values(rq)
	url := fmt.Sprintf("%v%v?%v", c.BaseURL, "events", v.Encode())
	eventsResponse := EventsResponse{}
	err := c.getResponse(url, &eventsResponse)
	if err != nil {
		return nil, err
	}
	return &eventsResponse, nil
}
