# myClubhouse API Client

## Instantiation
```golang
import "github.com/c-m-hunt/myclubhouse/apiclient"

c := apiclient.MakeClient(subdomain, apiToken, nil)
```

## Users
```golang
users := []data.User{}
cr, err := c.Users(nil, &users)
```
## User
```golang
id := 1
user, err := c.User(id)
```
## Events
```golang
events := []data.Event{}
cr, err = c.Events(nil, &events)
```
## Event
```golang
id := 1
event, err := c.Event(id)
```

## Request query
All client requests accept a request query as first argument which is in the following structure which will be passed through
```golang
type RequestQuery struct {
	View         string `url:"view"`
	Sort         string `url:"sort"`
	Filter       string `url:"filter"`
	SelectedPage int    `url:"selectedPage"`
	PageSize     int    `url:"pageSize"`
}
```

Example:
```golang
q := apiclient.RequestQuery{
  PageSize: 200,
}
users := []data.User{}
cr, err := c.Users(&q, &users)
```