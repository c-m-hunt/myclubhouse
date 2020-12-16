package apiclient

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/c-m-hunt/myclubhouse/data"
	"github.com/google/go-querystring/query"
)

type ResponsePagination struct {
	ItemCount    int
	PageCount    int
	PageSize     int
	SelectedPage int
}

// UsersResponse - Response back from the client
type UsersResponse struct {
	Users      []data.User
	Pagination ResponsePagination
}

type EventsResponse struct {
	Events     []data.Event
	Pagination ResponsePagination
}

// Request - Query used to request different variables
type RequestQuery struct {
	View         string `url:"view"`
	Sort         string `url:"sort"`
	Filter       string `url:"filter"`
	SelectedPage int    `url:"selectedPage"`
	PageSize     int    `url:"pageSize"`
}

func (c Client) getRequest(url string) *http.Request {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	req.Header.Set("X-ApiAccessToken", c.AccessToken)
	if err != nil {
		log.Fatal(err)
	}
	return req
}

// Users - Get users from the api client
func (c Client) Users(rq *RequestQuery) (*[]data.User, error) {
	ur, err := c.requestUsers(rq)
	return &ur.Users, err
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
	body, err := c.getResponse(url)
	if err != nil {
		return nil, err
	}
	jsonErr := json.Unmarshal(body, &usersResponse)
	if jsonErr != nil {
		return nil, jsonErr
	}
	return &usersResponse, nil
}

func (c Client) getResponse(url string) ([]byte, error) {
	req := c.getRequest(url)
	res, getErr := c.HTTPClient.Do(req)
	if getErr != nil {
		return nil, getErr
	}
	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		return nil, readErr
	}
	return body, nil
}

// RequestEvents - Requests events from the client
func (c Client) requestEvents(rq *RequestQuery) (*EventsResponse, error) {
	v, _ := query.Values(rq)
	url := fmt.Sprintf("%v%v?%v", c.BaseURL, "events", v.Encode())
	eventsResponse := EventsResponse{}
	body, err := c.getResponse(url)
	if err != nil {
		return nil, err
	}
	jsonErr := json.Unmarshal(body, &eventsResponse)
	if jsonErr != nil {
		return nil, jsonErr
	}
	return &eventsResponse, nil
}
