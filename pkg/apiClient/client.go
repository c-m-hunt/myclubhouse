package apiclient

import (
	"fmt"
	"net/http"
)

type Client struct {
	SubDomain   string
	AccessToken string
	BaseURL     string
	HTTPClient  http.Client
}

func NewClient(subDomain string, accessToken string) Client {
	baseURL := fmt.Sprintf("https://%v.myclubhouse.co.uk/api/v1/", subDomain)
	c := Client{
		SubDomain:   subDomain,
		AccessToken: accessToken,
		BaseURL:     baseURL,
	}
	c.HTTPClient = http.Client{
		//Timeout: time.Second * 2,
	}
	return c
}
