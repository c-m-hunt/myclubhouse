package apiclient

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/c-m-hunt/myclubhouse/pkg/data"
)

type UsersResponse struct {
	Users      []data.User
	Pagination struct {
		ItemCount    int
		PageCount    int
		PageSize     int
		SelectedPage int
	}
}

func (c Client) getRequest(url string) *http.Request {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	req.Header.Set("X-ApiAccessToken", c.AccessToken)
	if err != nil {
		log.Fatal(err)
	}
	return req
}

func (c Client) Users() (*[]data.User, error) {
	usersResponse, err := c.RequestUsers()
	return &usersResponse.Users, err
}

// Users Gets users from the service
func (c Client) RequestUsers() (*UsersResponse, error) {
	url := fmt.Sprintf("%v%v", c.BaseURL, "users")
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
