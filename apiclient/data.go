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

// UsersResponse - Response back from the client
type UsersResponse struct {
	Users      []data.User
	Pagination struct {
		ItemCount    int
		PageCount    int
		PageSize     int
		SelectedPage int
	}
}

// UsersQuery - Query used to request users
type UsersQuery struct {
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
func (c Client) Users(uq *UsersQuery) (*[]data.User, error) {
	usersResponse, err := c.requestUsers(uq)
	return &usersResponse.Users, err
}

// RequestUsers - Requests users from the client
func (c Client) requestUsers(uq *UsersQuery) (*UsersResponse, error) {
	v, _ := query.Values(uq)
	url := fmt.Sprintf("%v%v?%v", c.BaseURL, "users", v.Encode())
	req := c.getRequest(url)
	res, getErr := c.HTTPClient.Do(req)
	if getErr != nil {
		return nil, getErr
	}
	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		return nil, readErr
	}
	usersResponse := UsersResponse{}
	jsonErr := json.Unmarshal(body, &usersResponse)
	if jsonErr != nil {
		return nil, jsonErr
	}
	return &usersResponse, nil
}
