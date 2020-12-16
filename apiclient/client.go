package apiclient

import (
	"fmt"
	"net/http"
)

// Client - Api client struct
type Client struct {
	SubDomain   string
	AccessToken string
	BaseURL     string
	HTTPClient  *http.Client
}

// MakeClient - Generates an api client
func MakeClient(subDomain string, accessToken string, httpClient *http.Client) Client {
	baseURL := fmt.Sprintf("https://%v.myclubhouse.co.uk/api/v1/", subDomain)
	c := Client{
		SubDomain:   subDomain,
		AccessToken: accessToken,
		BaseURL:     baseURL,
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
