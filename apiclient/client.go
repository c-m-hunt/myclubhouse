package apiclient

import (
	"fmt"
	"net/http"
)

type Service interface {
	GetData(RequestQuery, *interface{})
}

type UsersService struct{}
type EventsService struct{}

// Client - Api client struct
type Client struct {
	SubDomain   string
	AccessToken string
	BaseURL     string
	HTTPClient  *http.Client
	UsersService
	EventsService
}

// MakeClient - Generates an api client
func MakeClient(subDomain string, accessToken string, httpClient *http.Client) Client {
	baseURL := fmt.Sprintf("https://%v.myclubhouse.co.uk/api/v1/", subDomain)
	c := Client{
		SubDomain:     subDomain,
		AccessToken:   accessToken,
		BaseURL:       baseURL,
		UsersService:  UsersService{},
		EventsService: EventsService{},
	}
	if httpClient == nil {
		c.HTTPClient = &http.Client{
			//Timeout: time.Second * 2,
			// Add general options here
		}
	} else {
		c.HTTPClient = httpClient
	}

	return c
}
