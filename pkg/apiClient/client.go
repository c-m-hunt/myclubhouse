package apiclient

import (
	"fmt"
	"net/http"
)

type Client struct {
	SubDomain   string
	AccessToken string
	BaseUrl     string
	HttpClient  http.Client
}

func NewClient(subDomain string, accessToken string) Client {
	c := Client{
		SubDomain:   subDomain,
		AccessToken: accessToken,
		BaseUrl:     fmt.Sprintf("https://%v.myclubhouse.co.uk/api/v1/", subDomain),
	}

	c.HttpClient = http.Client{
		//Timeout: time.Second * 2,
	}
	return c
}
