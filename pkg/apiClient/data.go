package apiclient

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type User struct {
	ID       int
	Username string
	Title    string
	Forename string
	Surname  string
}

type UsersResponse struct {
	Users      []User
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

// Users Gets users from the service
func (c Client) Users() UsersResponse {
	url := fmt.Sprintf("%v%v", c.BaseURL, "users")
	req := c.getRequest(url)
	res, getErr := c.HTTPClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}
	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}
	usersResponse := UsersResponse{}
	jsonErr := json.Unmarshal(body, &usersResponse)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	fmt.Print(usersResponse.Pagination)
	return usersResponse
}
