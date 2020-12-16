# myClubhouse API Client

## Instantiation
```golang
import "github.com/c-m-hunt/myclubhouse/apiclient"

c := apiclient.MakeClient(subdomain, apiToken)
```

## Users
```golang
q := apiclient.UsersQuery{
  PageSize: 200,
}
users, err := c.Users(&q)
```