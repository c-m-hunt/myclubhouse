# myClubhouse API Client

## Instantiation
```golang
import "github.com/c-m-hunt/myclubhouse/apiclient"

c := apiclient.MakeClient(subdomain, apiToken, nil)
```

## Users
```golang
q := apiclient.RequestQuery{
  PageSize: 200,
}
users, err := c.Users(&q)
```

## Events
```golang
q := apiclient.RequestQuery{
  PageSize: 200,
}
events, _ := c.Events(&q)
```