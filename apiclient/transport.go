package apiclient

import (
	"encoding/json"
	"log"
	"net/http"
)

// ResponsePagination - Struct detailing what data has been returned and is available
type ResponsePagination struct {
	ItemCount    int
	PageCount    int
	PageSize     int
	SelectedPage int
}

// ClientResponse - Response structure from any of the multiple requests
type ClientResponse struct {
	Data       interface{}
	Pagination ResponsePagination
}

// RequestQuery - Query used to request different variables
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

func (c Client) getResponse(url string, v interface{}) error {
	req := c.getRequest(url)
	res, getErr := c.HTTPClient.Do(req)
	if getErr != nil {
		return getErr
	}
	defer res.Body.Close()

	if err := json.NewDecoder(res.Body).Decode(&v); err != nil {
		return err
	}
	return nil
}
