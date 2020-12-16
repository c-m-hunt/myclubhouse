# MyClubHouse API Client

## Instantiation
```golang
import "github.com/c-m-hunt/myclubhouse/pkg/apiclient"

c := apiclient.MakeClient(subdomain, apiToken)
```

## Users
```golang
q := apiclient.UsersQuery{
  PageSize: 200,
}
users, _ := c.Users(&q)
```